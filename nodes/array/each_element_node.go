package array

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	eachElementArrayNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "EachElementArrayNode", FriendlyName: "Each Element In Array", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array"}}
	eachElementArrayNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Loop on all element in a array"}}
)

type EachElementArrayNode struct {
	block.NodeBase
}

func NewEachElementArrayNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(EachElementArrayNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewNodeParameter(node, "array", block.NodeParameterTypeEnumArray, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(array)

	each, err := block.NewDynamicNodeParameter(node, "each", block.NodeParameterTypeEnumNode, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(each)

	item, err := block.NewDynamicNodeParameter(node, "item", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(item)

	return node, nil
}

func (n *EachElementArrayNode) CanExecute() bool {
	return true
}

func (n *EachElementArrayNode) CanBeExecuted() bool {
	return true
}

func (n *EachElementArrayNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return eachElementArrayNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return eachElementArrayNodeGraphDescription
	default:
		return nil
	}
}

func (n *EachElementArrayNode) OnExecution(ctx context.Context, engine block.Engine) error {
	value := n.NodeData.InParameters.Get("array").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "array"}
	}

	array, ok := value.(*[]interface{})
	if !ok {
		return block.ErrInvalidParameter{Name: "array"}
	}

	each := n.Data().OutParameters.Get("each").Value
	if each == nil {
		return nil
	}

	eachNode, ok := each.(block.Node)
	if !ok {
		return nil
	}

	for _, item := range *array {
		n.Data().OutParameters.Get("item").Value = item
		engine.ExecuteNode(ctx, eachNode, nil)
	}

	return nil
}
