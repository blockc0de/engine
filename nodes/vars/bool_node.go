package vars

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	boolNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "BoolNode", FriendlyName: "Boolean", NodeType: attributes.NodeTypeEnumVariable, GroupName: "Base Variable", BlockLimitPerGraph: -1}}
	boolNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "In computer science, the Boolean data type is a data type that has one of two possible values (usually denoted true and false) which is intended to represent the two truth values of logic and Boolean algebra."}}
)

type BoolNode struct {
	block.NodeBase
}

func NewBoolNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(BoolNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumBool, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *BoolNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return boolNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return boolNodeGraphDescription
	default:
		return nil
	}
}
