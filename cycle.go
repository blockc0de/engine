package engine

import (
	"context"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/functions"
	"math/big"
	"reflect"
)

type GraphExecutionCycle struct {
	engine                          *Engine
	Timestamp                       int64
	StartNode                       block.StartNode
	ExecutedNodesInCycle            []block.Node
	Trace                           *block.GraphTrace
	StartNodeInstantiatedParameters block.NodeParameters
	CurrentFunctionContext          *functions.FunctionContext
}

func NewGraphExecutionCycle(engine *Engine,
	timestamp int64, startNode block.StartNode, parameters block.NodeParameters) *GraphExecutionCycle {

	cycle := GraphExecutionCycle{
		engine:               engine,
		Timestamp:            timestamp,
		StartNode:            startNode,
		ExecutedNodesInCycle: make([]block.Node, 0),
		Trace:                block.NewGraphTrace(),
	}

	cycle.addExecutedNode(startNode)
	cycle.StartNodeInstantiatedParameters = parameters
	if parameters == nil {
		cycle.StartNodeInstantiatedParameters = cycle.StartNode.Data().OutParameters
	}

	return &cycle
}

func (c *GraphExecutionCycle) Execute(ctx context.Context) {
	c.StartNode.Data().OutParameters = c.StartNodeInstantiatedParameters
	c.StartNode.BeginCycle(ctx, c.engine)
}

func (c *GraphExecutionCycle) GetCycleExecutedGasPrice() *big.Int {
	total := big.NewInt(0)
	for _, node := range c.ExecutedNodesInCycle {
		attrs := node.GetCustomAttributes(reflect.TypeOf(attributes.NodeGasConfiguration{}))
		if len(attrs) == 0 {
			continue
		}
		total = big.NewInt(0).Add(total, attrs[0].(attributes.NodeGasConfiguration).BlockGasPrice)
	}
	return total
}

func (c *GraphExecutionCycle) GetCycleMaxExecutionTime() int64 {
	var maxTimeout int64
	baseTime := int64(1000 * 60)
	for _, node := range c.engine.Graph.NodeList {
		if node.Data().CustomTimeout > maxTimeout {
			maxTimeout = node.Data().CustomTimeout
		}
	}
	return baseTime + maxTimeout
}

func (c *GraphExecutionCycle) addExecutedNode(node block.Node) *block.NodeTrace {
	c.ExecutedNodesInCycle = append(c.ExecutedNodesInCycle, node)
	return c.Trace.AppendTrace(node)
}
