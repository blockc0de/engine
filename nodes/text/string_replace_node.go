package text

import (
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringReplaceNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringReplaceNode", FriendlyName: "Replace String", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String", BlockLimitPerGraph: -1}}
	stringReplaceNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Split two string by a specific delimiter and return it as out parameter."}}
)

type StringReplaceNode struct {
	block.NodeBase
}

func NewStringReplaceNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringReplaceNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	original, err := block.NewNodeParameter(node, "original", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(original)

	toReplace, err := block.NewNodeParameter(node, "toReplace", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(toReplace)

	replaceText, err := block.NewNodeParameter(node, "replaceText", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(replaceText)

	outParameter, err := block.NewDynamicNodeParameter(node, "string", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, err
}

func (n *StringReplaceNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringReplaceNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringReplaceNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringReplaceNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		original := n.Data().InParameters.Get("original")
		toReplace := n.Data().InParameters.Get("toReplace")
		replaceText := n.Data().InParameters.Get("replaceText")

		var converter block.NodeParameterConverter
		originalAVal, ok := converter.ToString(original.ComputeValue())
		if !ok {
			return nil
		}

		toReplaceVal, ok := converter.ToString(toReplace.ComputeValue())
		if !ok {
			return nil
		}

		replaceTextVal, ok := converter.ToString(replaceText.ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(strings.ReplaceAll(originalAVal, toReplaceVal, replaceTextVal))
	}
	return value
}
