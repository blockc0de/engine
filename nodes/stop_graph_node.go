package nodes

import (
	"context"
	"github.com/graphlinq-go/engine/attributes"
	"github.com/graphlinq-go/engine/block"
	"math/big"
	"reflect"
)

var (
	stopGraphNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StopGraphNode", FriendlyName: "Stop Graph", NodeType: attributes.NodeTypeEnumEntryPoint, GroupName: "Common", BlockLimitPerGraph: -1}}
	stopGraphNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Stop the execution of the current graph"}}
	stopGraphNodeGasConfiguration = []interface{}{attributes.NodeGasConfiguration{BlockGasPrice: big.NewInt(0)}}
)

type StopGraphNode struct {
	block.NodeBase
}

func NewStopGraphNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StopGraphNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	return node, nil
}

func (n *StopGraphNode) CanBeExecuted() bool {
	return true
}

func (n *StopGraphNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stopGraphNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stopGraphNodeGraphDescription
	case reflect.TypeOf(attributes.NodeGasConfiguration{}):
		return stopGraphNodeGasConfiguration
	default:
		return nil
	}
}

func (n *StopGraphNode) OnExecution(ctx context.Context, executor block.NodeExecutor) error {
	executor.Stop()
	return nil
}
