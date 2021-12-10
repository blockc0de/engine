package text

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
	"strings"
)

var (
	stringTrimRightNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringTrimRightNode", FriendlyName: "Trim Right String", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String", BlockLimitPerGraph: -1}}
	stringTrimRightNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns a slice of the string with all trailing."}}
)

type StringTrimRightNode struct {
	block.NodeBase
}

func NewStringTrimRightNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringTrimRightNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	str, err := block.NewNodeParameter(node, "input", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(str)

	cutset, err := block.NewNodeParameter(node, "cutset", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(cutset)

	outParameter, err := block.NewDynamicNodeParameter(node, "string", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, err
}

func (n *StringTrimRightNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringTrimRightNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringTrimRightNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringTrimRightNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		str := n.Data().InParameters.Get("input")
		cutset := n.Data().InParameters.Get("cutset")

		var converter block.NodeParameterConverter
		stringVal, ok := converter.ToString(str.ComputeValue())
		if !ok {
			return nil
		}

		cursetVal, ok := converter.ToString(cutset.ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(strings.TrimRight(stringVal, cursetVal))
	}
	return value
}
