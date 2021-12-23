package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	mulNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "MultiplyNode", FriendlyName: "Multiply A * B", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	mulNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the multiplication of value of A * B (both sent in params) and return it as out parameter."}}
)

type MulNode struct {
	block.NodeBase
}

func NewMulNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(MulNode)
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

func (n *MulNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return mulNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return mulNodeGraphDescription
	default:
		return nil
	}
}

func (n *MulNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
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

		return block.NodeParameterDecimal{Decimal: a.Mul(b)}
	}
	return value
}
