package nodes

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stopGraphNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StopGraphNode", FriendlyName: "Stop Graph", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Common"}}
	stopGraphNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Stop the execution of the current graph"}}
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
	default:
		return nil
	}
}

func (n *StopGraphNode) OnExecution(ctx context.Context, engine block.Engine) error {
	engine.Stop()
	return nil
}
