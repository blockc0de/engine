package nodes

import (
	"context"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	onGraphStartNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnGraphStartNode", FriendlyName: "On Graph Start", NodeType: attributes.NodeTypeEnumEntryPoint, GroupName: "Common", BlockLimitPerGraph: -1}}
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

func (n *OnGraphStartNode) SetupEvent(ctx context.Context, executor block.NodeExecutor) error {
	executor.ExecuteNode(ctx, n, nil)
	return nil
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

func (n *OnGraphStartNode) OnExecution(context.Context, block.NodeExecutor) error {
	return nil
}
