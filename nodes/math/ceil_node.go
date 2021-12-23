package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	ceilNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CeilNode", FriendlyName: "Ceiling", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	ceilNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Ceiling the value"}}
)

type CeilNode struct {
	block.NodeBase
}

func NewCeilNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CeilNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	number, err := block.NewNodeParameter(node, "number", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(number)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *CeilNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return ceilNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return ceilNodeGraphDescription
	default:
		return nil
	}
}

func (n *CeilNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		var converter block.NodeParameterConverter
		number, ok := converter.ToDecimal(n.Data().InParameters.Get("number").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterDecimal{Decimal: number.Ceil()}
	}
	return value
}
