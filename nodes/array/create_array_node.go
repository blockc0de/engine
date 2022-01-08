package array

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	createArrayNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CreateArrayNode", FriendlyName: "Create Array", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array"}}
	createArrayNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "An array can store multiple variable in it"}}
)

type CreateArrayNode struct {
	block.NodeBase
}

func NewCreateArrayNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CreateArrayNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewDynamicNodeParameter(node, "array", block.NodeParameterTypeEnumArray, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(array)

	return node, nil
}

func (n *CreateArrayNode) CanExecute() bool {
	return true
}

func (n *CreateArrayNode) CanBeExecuted() bool {
	return true
}

func (n *CreateArrayNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return createArrayNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return createArrayNodeGraphDescription
	default:
		return nil
	}
}

func (n *CreateArrayNode) OnExecution(context.Context, block.NodeScheduler) error {
	array := make([]interface{}, 0)
	n.NodeData.OutParameters.Get("array").Value = &array
	return nil
}
