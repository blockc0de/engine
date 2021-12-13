package array

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	getElementAtIndexNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetElementAtIndexNode", FriendlyName: "Get Array Element At Index", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array", BlockLimitPerGraph: -1}}
	getElementAtIndexNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get a element from a array at a specific index"}}
)

type GetElementAtIndexNode struct {
	block.NodeBase
}

func NewGetElementAtIndexNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetElementAtIndexNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewNodeParameter(node, "array", block.NodeParameterTypeEnumArray, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(array)

	index, err := block.NewNodeParameter(node, "index", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(index)

	element, err := block.NewDynamicNodeParameter(node, "element", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(element)

	return node, nil
}

func (n *GetElementAtIndexNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getElementAtIndexNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getElementAtIndexNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetElementAtIndexNode) ComputeParameterValue(id string, value interface{}) interface{} {
	if id == n.Data().OutParameters.Get("element").Id {
		value := n.NodeData.InParameters.Get("array").ComputeValue()
		if value == nil {
			return nil
		}

		array, ok := value.(*[]interface{})
		if !ok {
			return nil
		}

		var converter block.NodeParameterConverter
		index, ok := converter.ToDecimal(n.NodeData.InParameters.Get("index").ComputeValue())
		if !ok {
			return nil
		}

		n := int(index.IntPart())
		if n >= 0 && n < len(*array) {
			return (*array)[n]
		}
		return nil
	}
	return value
}
