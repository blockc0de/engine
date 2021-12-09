package math

import (
	"github.com/graphlinq-go/engine/attributes"
	"github.com/graphlinq-go/engine/block"
	"reflect"
)

var (
	subNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "SubtractNode", FriendlyName: "Substract A - B", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	subNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the substraction of value of A - B (both sent in params) and return it as out parameter."}}
)

type SubNode struct {
	block.NodeBase
}

func NewSubNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(SubNode)
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

func (n *SubNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return subNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return subNodeGraphDescription
	default:
		return nil
	}
}

func (n *SubNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		a := n.Data().InParameters.Get("a")
		b := n.Data().InParameters.Get("b")

		var converter block.NodeParameterConverter
		aVal, ok := converter.ToDecimal(a.ComputeValue())
		if !ok {
			return nil
		}

		bVal, ok := converter.ToDecimal(b.ComputeValue())
		if !ok {
			return nil
		}
		bVal.Ceil()
		return block.NodeParameterDecimal{Decimal: aVal.Sub(bVal)}
	}
	return value
}
