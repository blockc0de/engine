package math

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	truncNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "TruncateNode", FriendlyName: "Truncate", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	truncNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Truncates off digits from the number, without rounding."}}
)

type TruncNode struct {
	block.NodeBase
}

func NewTruncNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(TruncNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	number, err := block.NewNodeParameter(node, "number", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(number)

	precision, err := block.NewNodeParameter(node, "precision", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(precision)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *TruncNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return truncNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return truncNodeGraphDescription
	default:
		return nil
	}
}

func (n *TruncNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		number := n.Data().InParameters.Get("number")
		precision := n.Data().InParameters.Get("precision")

		var converter block.NodeParameterConverter
		numberVal, ok := converter.ToDecimal(number.ComputeValue())
		if !ok {
			return nil
		}

		precisionVal, ok := converter.ToDecimal(precision.ComputeValue())
		if !ok {
			return nil
		}
		return block.NodeParameterDecimal{Decimal: numberVal.Truncate(int32(precisionVal.IntPart()))}
	}
	return value
}
