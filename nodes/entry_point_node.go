package nodes

import (
	"context"
	"github.com/graphlinq-go/engine/attributes"
	"github.com/graphlinq-go/engine/block"
	"reflect"
)

var (
	entryPointNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "EntryPointNode", FriendlyName: "Entry Point", NodeType: attributes.NodeTypeEnumEntryPoint, GroupName: "Common", BlockLimitPerGraph: -1}}
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

func (n *EntryPointNode) OnExecution(context.Context, block.NodeExecutor) error {
	return nil
}
