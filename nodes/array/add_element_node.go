package array

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	addElementNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddArrayElementNode", FriendlyName: "Add Array Element", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array", BlockLimitPerGraph: -1}}
	addElementNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Add a element to the array"}}
)

type AddElementNode struct {
	block.NodeBase
}

func NewAddElementNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AddElementNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewNodeParameter(node, "array", block.NodeParameterTypeEnumArray, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(array)

	element, err := block.NewNodeParameter(node, "element", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(element)

	return node, nil
}

func (n *AddElementNode) CanExecute() bool {
	return true
}

func (n *AddElementNode) CanBeExecuted() bool {
	return true
}

func (n *AddElementNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return addElementNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return addElementNodeGraphDescription
	default:
		return nil
	}
}

func (n *AddElementNode) OnExecution(context.Context, block.NodeScheduler) error {
	value := n.NodeData.InParameters.Get("array").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "array"}
	}

	array, ok := value.(*[]interface{})
	if !ok {
		return block.ErrInvalidParameter{Name: "array"}
	}

	element := n.NodeData.InParameters.Get("element").ComputeValue()
	if element == nil {
		return block.ErrInvalidParameter{Name: "element"}
	}

	*array = append(*array, element)

	return nil
}
