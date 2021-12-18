package nodes

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	onGraphStartNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnGraphStartNode", FriendlyName: "On Graph Start", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Common", BlockLimitPerGraph: -1}}
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

func (n *OnGraphStartNode) SetupEvent(scheduler block.NodeScheduler) error {
	scheduler.AddCycle(n, nil)
	return nil
}

func (n *OnGraphStartNode) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)
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

func (n *OnGraphStartNode) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *OnGraphStartNode) OnStop() error {
	return nil
}
