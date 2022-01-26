package nodes

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	onGraphStartNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnGraphStartNode", FriendlyName: "On Graph Start", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Common"}}
	onGraphStartNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "This event is called when the graph start, usefull for initialize variables"}}
)

type OnGraphStartNode struct {
	block.NodeBase
}

func NewOnGraphStartNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(OnGraphStartNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true
	return node, nil
}

func (n *OnGraphStartNode) CanExecute() bool {
	return true
}

func (n *OnGraphStartNode) SetupEvent(engine block.Engine) error {
	engine.AddCycle(n, nil)
	return nil
}

func (n *OnGraphStartNode) BeginCycle(ctx context.Context, engine block.Engine) {
	engine.NextNode(ctx, n)
}

func (n *OnGraphStartNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return onGraphStartNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return onGraphStartNodeGraphDescription
	default:
		return nil
	}
}

func (n *OnGraphStartNode) OnExecution(context.Context, block.Engine) error {
	return nil
}

func (n *OnGraphStartNode) OnStop() error {
	return nil
}
