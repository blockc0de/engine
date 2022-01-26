package array

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	clearArrayNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "ClearArrayNode", FriendlyName: "Clear Array", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array"}}
	clearArrayNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Clear all elements in an array"}}
)

type ClearArrayNode struct {
	block.NodeBase
}

func NewClearArrayNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(ClearArrayNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewNodeParameter(node, "array", block.NodeParameterTypeEnumArray, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(array)

	return node, nil
}

func (n *ClearArrayNode) CanExecute() bool {
	return true
}

func (n *ClearArrayNode) CanBeExecuted() bool {
	return true
}

func (n *ClearArrayNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return clearArrayNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return clearArrayNodeGraphDescription
	default:
		return nil
	}
}

func (n *ClearArrayNode) OnExecution(context.Context, block.Engine) error {
	value := n.NodeData.InParameters.Get("array").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "array"}
	}

	array, ok := value.(*[]interface{})
	if !ok {
		return block.ErrInvalidParameter{Name: "array"}
	}

	*array = (*array)[:0]

	return nil
}
