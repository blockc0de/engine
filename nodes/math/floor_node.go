package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	floorNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "FloorNode", FriendlyName: "Floor", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	floorNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Floor the value"}}
)

type FloorNode struct {
	block.NodeBase
}

func NewFloorNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(FloorNode)
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

func (n *FloorNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return floorNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return floorNodeGraphDescription
	default:
		return nil
	}
}

func (n *FloorNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		var converter block.NodeParameterConverter
		number, ok := converter.ToDecimal(n.Data().InParameters.Get("number").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterDecimal{Decimal: number.Floor()}
	}
	return value
}
