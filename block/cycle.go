package block

import (
	"context"
	"math/big"
	"reflect"

	"github.com/blockc0de/engine/attributes"
)

type LocalStorage struct {
	m map[string]interface{}
}

func (s LocalStorage) Get(key string) interface{} {
	val, ok := s.m[key]
	if !ok {
		return nil
	}
	return val
}

func (s LocalStorage) Add(key string, val interface{}) {
	s.m[key] = val
}

type GraphExecutionCycle struct {
	engine                          Engine
	Timestamp                       int64
	StartNode                       StartNode
	ExecutedNodesInCycle            []Node
	Trace                           *GraphTrace
	LocalStorage                    LocalStorage
	StartNodeInstantiatedParameters NodeParameters
}

func NewGraphExecutionCycle(engine Engine,
	timestamp int64, startNode StartNode, parameters NodeParameters) *GraphExecutionCycle {

	cycle := GraphExecutionCycle{
		engine:               engine,
		Timestamp:            timestamp,
		StartNode:            startNode,
		ExecutedNodesInCycle: make([]Node, 0),
		LocalStorage:         LocalStorage{m: map[string]interface{}{}},
		Trace:                NewGraphTrace(),
	}

	cycle.AddExecutedNode(startNode)
	cycle.StartNodeInstantiatedParameters = parameters
	if parameters == nil {
		cycle.StartNodeInstantiatedParameters = cycle.StartNode.Data().OutParameters
	}

	return &cycle
}

func (c *GraphExecutionCycle) Execute(ctx context.Context) {
	for idx := range c.StartNode.Data().OutParameters {
		if idx >= len(c.StartNodeInstantiatedParameters) {
			break
		}
		*c.StartNode.Data().OutParameters[idx] = *c.StartNodeInstantiatedParameters[idx]
	}
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

func (c *GraphExecutionCycle) AddExecutedNode(node Node) *NodeTrace {
	c.ExecutedNodesInCycle = append(c.ExecutedNodesInCycle, node)
	return c.Trace.AppendTrace(node)
}
