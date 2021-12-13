package array

import (
	"reflect"

	"github.com/shopspring/decimal"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	getArraySizeNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetArraySizeNode", FriendlyName: "Get Array Size", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Array", BlockLimitPerGraph: -1}}
	getArraySizeNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get the size of an array"}}
)

type GetArraySizeNode struct {
	block.NodeBase
}

func NewGetArraySizeNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetArraySizeNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewNodeParameter(node, "array", block.NodeParameterTypeEnumArray, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(array)

	size, err := block.NewDynamicNodeParameter(node, "size", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(size)

	return node, nil
}

func (n *GetArraySizeNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getArraySizeNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getArraySizeNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetArraySizeNode) ComputeParameterValue(id string, value interface{}) interface{} {
	if id == n.Data().OutParameters.Get("size").Id {
		value := n.NodeData.InParameters.Get("array").ComputeValue()
		if value == nil {
			return nil
		}

		array, ok := value.(*[]interface{})
		if !ok {
			return nil
		}

		return block.NodeParameterDecimal{Decimal: decimal.NewFromInt(int64(len(*array)))}
	}
	return value
}
