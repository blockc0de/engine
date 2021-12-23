package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/shopspring/decimal"
)

var (
	divNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "DivideNode", FriendlyName: "Divide A / B", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	divNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the division of value of A / B (both sent in params) and return it as out parameter, zero division are forbidden."}}
)

type DivNode struct {
	block.NodeBase
}

func NewDivNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(DivNode)
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

func (n *DivNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return divNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return divNodeGraphDescription
	default:
		return nil
	}
}

func (n *DivNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
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

		if a.IsZero() || b.IsZero() {
			return block.NodeParameterDecimal{Decimal: decimal.Zero}
		}
		return block.NodeParameterDecimal{Decimal: a.Div(b)}
	}
	return value
}
