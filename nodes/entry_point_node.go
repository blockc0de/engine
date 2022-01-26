package nodes

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	entryPointNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "EntryPointNode", FriendlyName: "Entry Point", NodeType: attributes.NodeTypeEnumEntryPoint, GroupName: "Common"}}
	entryPointNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Basic Graph entry point, start the execution of a graph"}}
)

type EntryPointNode struct {
	block.NodeBase
}

func NewEntryPointNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(EntryPointNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	return node, nil
}

func (n *EntryPointNode) CanExecute() bool {
	return true
}

func (n *EntryPointNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return entryPointNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return entryPointNodeGraphDescription
	default:
		return nil
	}
}

func (n *EntryPointNode) BeginCycle(ctx context.Context, engine block.Engine) {
	engine.NextNode(ctx, n)
}

func (n *EntryPointNode) OnExecution(context.Context, block.Engine) error {
	return nil
}
