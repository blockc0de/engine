package engine

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes"
	"github.com/blockc0de/engine/nodes/functions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
)

var (
	onGraphStartNode block.Node
)

func init() {
	var err error
	onGraphStartNode, err = nodes.NewOnGraphStartNode("", nil)
	if err != nil {
		panic(err)
	}
}

type Event struct {
	CycleCost func(cost *big.Int)
	AppendLog func(msgType string, message string)
}

// EventNodeSlice attaches the methods of Interface to []block.EventNode, sorting in increasing order.
type EventNodeSlice []block.EventNode

func (x EventNodeSlice) Len() int { return len(x) }

func (x EventNodeSlice) Less(i, j int) bool {
	iIsOnGraphStartNode := x[i].Data().FriendlyName == onGraphStartNode.Data().FriendlyName
	jIsOnGraphStartNode := x[j].Data().FriendlyName == onGraphStartNode.Data().FriendlyName
	if iIsOnGraphStartNode && jIsOnGraphStartNode {
		return i < j
	}
	if iIsOnGraphStartNode {
		return true
	}
	if jIsOnGraphStartNode {
		return false
	}
	return i < j
}

func (x EventNodeSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x EventNodeSlice) Sort() { sort.Sort(x) }

type Engine struct {
	Graph              *block.Graph
	ExecutedNodes      []block.Node
	running            bool
	stopping           bool
	client             redis.Cmdable
	owner              common.Address
	event              Event
	context            context.Context
	cancel             context.CancelFunc
	countdown          int
	currentCycle       *block.GraphExecutionCycle
	pendingCyclesQueue chan *block.GraphExecutionCycle
}

func NewEngine(graph *block.Graph, owner common.Address, client redis.Cmdable, event Event) *Engine {
	engine := Engine{
		Graph:         graph,
		event:         event,
		client:        client,
		owner:         owner,
		ExecutedNodes: make([]block.Node, 0),
	}
	return &engine
}

func (e *Engine) Run(ctx context.Context) error {
	if e.running {
		return errors.New("engine is running")
	}

	e.running = true
	e.stopping = false
	e.ExecutedNodes = e.ExecutedNodes[:0]
	e.context, e.cancel = context.WithCancel(ctx)
	e.pendingCyclesQueue = make(chan *block.GraphExecutionCycle, 128)
	maxExecutionTime := time.Second * time.Duration(e.Graph.GetMaxExecutionTime())

	e.startNodes()

loop:
	for {
		select {
		case <-ctx.Done():
			e.Stop()
		case cycle, ok := <-e.pendingCyclesQueue:
			if !ok {
				break loop
			}
			if e.stopping {
				continue
			}

			lastCycleAt := cycle.StartNode.Data().LastCycleAt
			nodeCycleLimit := cycle.StartNode.Data().NodeCycleLimit
			if lastCycleAt+nodeCycleLimit > time.Now().UnixMilli() {
				e.AppendLog("warn", "Reach node cycle limit'"+cycle.StartNode.Data().FriendlyName+"'")
				continue
			}

			e.currentCycle = cycle
			e.Graph.CurrentCycle = cycle
			cycle.StartNode.Data().LastCycleAt = time.Now().UnixMilli()

			c, cancel := context.WithTimeout(e.context, maxExecutionTime)
			cycle.Execute(c)
			if c.Err() != nil && c.Err() != context.Canceled {
				e.AppendLog("error", "Timeout occurred on last cycle from graph hash: "+e.Graph.Hash)
			}
			cancel()

			if e.event.CycleCost != nil {
				e.event.CycleCost(cycle.GetCycleExecutedGasPrice())
			}

			e.currentCycle = nil
			e.Graph.CurrentCycle = nil

			if e.countdown > 0 {
				e.countdown -= 1
				if e.countdown == 0 {
					e.Stop()
				}
			}
		}
	}

	e.stopNodes()

	e.cancel = nil
	e.context = nil
	e.running = false
	e.stopping = false

	return nil
}

func (e *Engine) Stop() {
	if !e.running || e.stopping {
		return
	}

	e.stopping = true

	e.cancel()
	close(e.pendingCyclesQueue)
}

func (e *Engine) AppendLog(msgType string, message string) {
	if e.event.AppendLog != nil {
		e.event.AppendLog(msgType, message)
	}
}

func (e *Engine) AddCycle(startNode block.StartNode, parameters block.NodeParameters) {
	if !e.running || e.stopping {
		return
	}

	cycle := block.NewGraphExecutionCycle(e, time.Now().Unix(), startNode, parameters)
	e.pendingCyclesQueue <- cycle
}

func (e *Engine) CurrentCycle() *block.GraphExecutionCycle {
	return e.currentCycle
}

func (e *Engine) NextNode(ctx context.Context, node block.Node) bool {
	if node.Data().OutNode != nil {
		return e.ExecuteNode(ctx, node.Data().OutNode, node)
	}
	return false
}

func (e *Engine) ExecuteNode(ctx context.Context, node block.Node, executedFromNode block.Node) bool {
	if !e.running || e.stopping || e.currentCycle == nil {
		return false
	}

	node.Data().LastExecutionFrom = executedFromNode

	executableNode, ok := node.(block.ExecutableNode)
	if !ok {
		return false
	}

	traceItem := e.currentCycle.AddExecutedNode(node)

	if node.Data().NodeType != reflect.TypeOf(new(nodes.EntryPointNode)).String() &&
		node.Data().NodeType != reflect.TypeOf(new(functions.FunctionNode)).String() {
		if !node.CanBeExecuted() {
			return false
		}
	}

	startTime := time.Now()
	executableNode.Data().CurrentTraceItem = traceItem

	err := executableNode.OnExecution(ctx, e)
	if err != nil {
		e.AppendLog("error", fmt.Sprintf("[%s] Error on node execution, reason: %s",
			executableNode.Data().FriendlyName, err.Error()))

		if executableNode.Data().CurrentTraceItem != nil {
			executableNode.Data().CurrentTraceItem.ExecutionError = err
		}

		e.Stop()
		return false
	}

	elapsedTime := time.Since(startTime)
	executableNode.Data().CurrentTraceItem = nil
	if traceItem != nil {
		traceItem.ExecutionTime = elapsedTime.Milliseconds()
	}

	return e.NextNode(ctx, node)
}

func (e *Engine) startNodes() {
	var numberOfEventNode int
	var numberOfConnectorNode int
	for _, node := range e.Graph.NodeList {
		if node.Data().IsEventNode {
			numberOfEventNode += 1
		}
		if node.Data().NodeBlockType == attributes.NodeTypeEnumConnector {
			numberOfConnectorNode += 1
		}
	}

	entryPointNode := e.Graph.GetFirstEntryPointNode()
	if numberOfEventNode == 0 {
		e.countdown = numberOfConnectorNode
		if entryPointNode != nil {
			e.countdown += 1
		}
	}

	// Init storage
	for _, node := range e.Graph.NodeList {
		storageNode, ok := node.(block.StorageNode)
		if !ok {
			continue
		}

		if err := storageNode.SetupDatabase(e.owner.String(), e.client); err != nil {
			e.AppendLog("error", "Can't setup the database: "+node.Data().FriendlyName+", "+err.Error())
			e.Stop()
			return
		}
	}

	// Init connectors
	for _, node := range e.Graph.NodeList {
		if node.Data().NodeBlockType == attributes.NodeTypeEnumConnector {
			connectorNode, ok := node.(block.ConnectorNode)
			if !ok {
				return
			}

			if err := connectorNode.SetupConnector(e); err != nil {
				e.AppendLog("error", "Can't setup the connector: "+node.Data().FriendlyName+", "+err.Error())
				e.Stop()
				return
			}
		}
	}

	// Setup event
	eventNodes := e.Graph.GetEventNodes()
	EventNodeSlice(eventNodes).Sort()
	for _, eventNode := range eventNodes {
		if err := eventNode.SetupEvent(e); err != nil {
			e.AppendLog("error", "Can't setup the event: "+eventNode.Data().FriendlyName+", "+err.Error())
			e.Stop()
			return
		}
	}

	// Execute entry point node
	if entryPointNode != nil {
		e.AddCycle(entryPointNode.(*nodes.EntryPointNode), nil)
	} else if numberOfEventNode == 0 && numberOfConnectorNode == 0 {
		e.Stop()
	}
}

func (e *Engine) stopNodes() {
	for _, node := range e.Graph.NodeList {
		if node.Data().NodeBlockType == attributes.NodeTypeEnumEvent {
			eventNode, ok := node.(block.EventNode)
			if !ok {
				continue
			}

			if err := eventNode.OnStop(); err != nil {
				e.AppendLog("error", "Can't release the connector "+eventNode.Data().FriendlyName+": "+err.Error())
			}
		}
	}

	for _, node := range e.Graph.NodeList {
		if node.Data().NodeBlockType == attributes.NodeTypeEnumConnector {
			connectorNode, ok := node.(block.ConnectorNode)
			if !ok {
				continue
			}

			if err := connectorNode.OnStop(); err != nil {
				e.AppendLog("error", "Can't release the event "+connectorNode.Data().FriendlyName+": "+err.Error())
			}
		}
	}

	e.AppendLog("warn", "Stop requested for graph hash: "+e.Graph.Hash)
}
