package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/shopspring/decimal"
)

var (
	percentageDiffNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "PercentageDiffNode", FriendlyName: "Percentage Difference", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	percentageDiffNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the percentage difference of value of A from B (both sent in params) and return difference as out parameter."}}
)

type PercentageDiffNode struct {
	block.NodeBase
}

func NewPercentageDiffNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(PercentageDiffNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	a, err := block.NewNodeParameter(node, "a", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(a)

	b, err := block.NewNodeParameter(node, "b", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(b)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *PercentageDiffNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return percentageDiffNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return percentageDiffNodeGraphDescription
	default:
		return nil
	}
}

func (n *PercentageDiffNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		var converter block.NodeParameterConverter
		a, ok := converter.ToDecimal(n.Data().InParameters.Get("a").ComputeValue())
		if !ok {
			return nil
		}

		b, ok := converter.ToDecimal(n.Data().InParameters.Get("b").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterDecimal{
			Decimal: b.Sub(a).Div(a.Abs()).Mul(decimal.NewFromInt(100))}
	}
	return value
}
