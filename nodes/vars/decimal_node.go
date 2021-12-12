package vars

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	decimalNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "DecimalNode", FriendlyName: "Decimal", NodeType: attributes.NodeTypeEnumVariable, GroupName: "Base Variable", BlockLimitPerGraph: -1}}
	decimalNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "The decimal data type is an exact numeric data type defined by its precision (total number of digits) and scale (number of digits to the right of the decimal point). ... Scale can be 0 (no digits to the right of the decimal point)."}}
)

type DecimalNode struct {
	block.NodeBase
}

func NewDecimalNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(DecimalNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *DecimalNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return decimalNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return decimalNodeGraphDescription
	default:
		return nil
	}
}
