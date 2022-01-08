package text

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringConcatNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringConcatNode", FriendlyName: "Concat String", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String"}}
	stringConcatNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Split two string by a specific delimiter and return it as out parameter."}}
)

type StringConcatNode struct {
	block.NodeBase
}

func NewStringConcatNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringConcatNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	stringA, err := block.NewNodeParameter(node, "stringA", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(stringA)

	stringB, err := block.NewNodeParameter(node, "stringB", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(stringB)

	delimiter, err := block.NewNodeParameter(node, "delimiter", block.NodeParameterTypeEnumString, true, block.NodeParameterString(""))
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(delimiter)

	outParameter, err := block.NewDynamicNodeParameter(node, "string", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, err
}

func (n *StringConcatNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringConcatNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringConcatNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringConcatNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		var converter block.NodeParameterConverter
		stringA, ok := converter.ToString(n.Data().InParameters.Get("stringA").ComputeValue())
		if !ok {
			return nil
		}

		stringB, ok := converter.ToString(n.Data().InParameters.Get("stringB").ComputeValue())
		if !ok {
			return nil
		}

		delimiter, ok := converter.ToString(n.Data().InParameters.Get("delimiter").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(stringA + delimiter + stringB)
	}
	return value
}
