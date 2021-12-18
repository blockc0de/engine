package math

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	addNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddNode", FriendlyName: "Add A + B", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Math", BlockLimitPerGraph: -1}}
	addNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Calculate the value of A + B (both sent in params) and return it as out parameter."}}
)

type AddNode struct {
	block.NodeBase
}

func NewAddNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AddNode)
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

func (n *AddNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return addNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return addNodeGraphDescription
	default:
		return nil
	}
}

func (n *AddNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
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

		return block.NodeParameterDecimal{Decimal: aVal.Add(bVal)}
	}
	return value
}
