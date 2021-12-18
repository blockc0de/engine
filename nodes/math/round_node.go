package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	roundNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "RoundNode", FriendlyName: "Round", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	roundNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Round the value"}}
)

type RoundNode struct {
	block.NodeBase
}

func NewRoundNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(RoundNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	number, err := block.NewNodeParameter(node, "number", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(number)

	places, err := block.NewNodeParameter(node, "places", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(places)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *RoundNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return roundNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return roundNodeGraphDescription
	default:
		return nil
	}
}

func (n *RoundNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		number := n.Data().InParameters.Get("number")
		places := n.Data().InParameters.Get("places")

		var converter block.NodeParameterConverter
		numberVal, ok := converter.ToDecimal(number.ComputeValue())
		if !ok {
			return nil
		}

		placesVal, ok := converter.ToDecimal(places.ComputeValue())
		if !ok {
			return nil
		}
		return block.NodeParameterDecimal{Decimal: numberVal.Round(int32(placesVal.IntPart()))}
	}
	return value
}
