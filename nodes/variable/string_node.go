package variable

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	stringNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringNode", FriendlyName: "String", NodeType: attributes.NodeTypeEnumVariable, GroupName: "Base Variable", BlockLimitPerGraph: -1}}
	stringNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "A string is a data type used in programming, such as an integer and floating point unit, but is used to represent text rather than numbers. It is comprised of a set of characters that can also contain spaces and numbers."}}
)

type StringNode struct {
	block.NodeBase
}

func NewStringNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumString, true, block.NodeParameterString(""))
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *StringNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringNodeGraphDescription
	default:
		return nil
	}
}
