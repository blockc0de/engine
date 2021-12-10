package engine

import (
	"context"
	"errors"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes"
	"github.com/blockc0de/engine/nodes/functions"
	"go.uber.org/zap"
	"reflect"
	"time"
)

type Engine struct {
	Graph         *block.Graph
	ExecutedNodes []block.Node
	Trace         block.GraphTrace
	running       bool
	logger        *zap.Logger
	appendLog     func(string, string)
	userData      map[string]interface{}
}

func NewEngine(graph *block.Graph, logger *zap.Logger, appendLog func(string, string)) *Engine {
	cycle := new(Engine)
	cycle.Graph = graph
	cycle.Trace = block.NewGraphTrace()
	cycle.ExecutedNodes = make([]block.Node, 0)

	cycle.appendLog = appendLog
	cycle.logger = logger.Named("Engine")
	return cycle
}

func (c *Engine) Run(ctx context.Context) error {
	if c.running {
		return errors.New("already running")
	}

	c.running = true

	for _, node := range c.Graph.NodeList {
		if node.Data().NodeBlockType == attributes.NodeTypeEnumConnector {
			connectorNode, ok := node.(block.ConnectorNode)
			if !ok {
				return errors.New("non connector node")
			}
			if err := connectorNode.SetupConnector(ctx, c); err != nil {
				c.appendLog("error", "Can't setup the connector : "+node.Data().FriendlyName+", "+err.Error())
				c.Stop()
				return err
			}
		}
	}

	eventNodes := c.Graph.GetEventNodes()
	for _, node := range eventNodes {
		if onGraphStartNode, ok := node.(*nodes.OnGraphStartNode); ok {
			if err := onGraphStartNode.SetupEvent(ctx, c); err != nil {
				c.appendLog("error", "Can't setup the event : "+node.Data().FriendlyName+", "+err.Error())
				c.Stop()
				return err
			}
		}
	}

	entryPointNode := c.Graph.GetFirstEntryPointNode()
	if entryPointNode == nil {
		return errors.New("entry point node not found")
	}

	c.ExecuteNode(ctx, entryPointNode, nil)
	return nil
}

func (c *Engine) Stop() {
	if !c.running {
		return
	}
}

func (c *Engine) NextNode(ctx context.Context, node block.Node) bool {
	if node.Data().OutNode != nil {
		return c.ExecuteNode(ctx, node.Data().OutNode, node)
	}
	return false
}

func (c *Engine) ExecuteNode(ctx context.Context, node block.Node, executedFromNode block.Node) bool {
	if !c.running {
		return false
	}

	node.Data().LastExecutionFrom = executedFromNode

	executableNode, ok := node.(block.ExecutableNode)
	if !ok {
		return false
	}

	traceItem := c.addExecutedNode(node)

	if node.Data().NodeType != reflect.TypeOf(new(nodes.EntryPointNode)).String() &&
		node.Data().NodeType != reflect.TypeOf(new(functions.FunctionNode)).String() {
		if !node.CanBeExecuted() {
			return false
		}
	}

	startTime := time.Now()
	executableNode.Data().CurrentTraceItem = traceItem

	err := executableNode.OnExecution(ctx, c)
	if err != nil {
		c.logger.Error("Error on node execution '"+node.Data().FriendlyName+"'", zap.Error(err))

		if executableNode.Data().CurrentTraceItem != nil {
			executableNode.Data().CurrentTraceItem.ExecutionError = err
		}
	}

	elapsedTime := time.Now().Sub(startTime)
	executableNode.Data().CurrentTraceItem = nil
	if traceItem != nil {
		traceItem.ExecutionTime = elapsedTime.Milliseconds()
	}

	return c.NextNode(ctx, node)
}

func (c *Engine) AppendLog(msgType string, message string) {
	if c.appendLog != nil {
		c.appendLog(msgType, message)
	}
}

func (c *Engine) GetUserData(key string) interface{} {
	data, _ := c.userData[key]
	return data
}

func (c *Engine) SetUserData(key string, data interface{}) {
	c.userData[key] = data
}

func (c *Engine) addExecutedNode(node block.Node) *block.NodeTrace {
	c.ExecutedNodes = append(c.ExecutedNodes, node)
	return c.Trace.AppendTrace(node)
}
