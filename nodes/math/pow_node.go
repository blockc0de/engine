package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/cockroachdb/apd"
	"github.com/shopspring/decimal"
)

var (
	powNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "PowerNode", FriendlyName: "Power A ^ B", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math"}}
	powNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the power raised to the base value. It takes two arguments. It returns the power raised to the base value."}}
)

type PowNode struct {
	block.NodeBase
}

func NewPowNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(PowNode)
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

func (n *PowNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return powNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return powNodeGraphDescription
	default:
		return nil
	}
}

func (n *PowNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
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

		d, _, _ := apd.NewFromString("0")
		x, _, _ := apd.NewFromString(a.String())
		y, _, _ := apd.NewFromString(b.String())
		if _, err := apd.BaseContext.Pow(d, x, y); err != nil {
			return nil
		}

		result, _ := decimal.NewFromString(d.String())
		return block.NodeParameterDecimal{Decimal: result}
	}
	return value
}
