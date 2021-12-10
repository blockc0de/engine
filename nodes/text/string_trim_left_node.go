package text

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
	"strings"
)

var (
	stringTrimLeftNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringTrimLeftNode", FriendlyName: "Trim Left String", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String", BlockLimitPerGraph: -1}}
	stringTrimLeftNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns a slice of the string with all leading."}}
)

type StringTrimLeftNode struct {
	block.NodeBase
}

func NewStringTrimLeftNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringTrimLeftNode)
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

func (n *StringTrimLeftNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringTrimLeftNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringTrimLeftNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringTrimLeftNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
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

		return block.NodeParameterString(strings.TrimLeft(stringVal, cursetVal))
	}
	return value
}
