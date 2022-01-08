package text

import (
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringToLowerNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringToLowerNode", FriendlyName: "String To Lower", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String"}}
	stringToLowerNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns s with all Unicode letters mapped to their lower case."}}
)

type StringToLowerNode struct {
	block.NodeBase
}

func NewStringToLowerNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringToLowerNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	str, err := block.NewNodeParameter(node, "input", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(str)

	outParameter, err := block.NewDynamicNodeParameter(node, "string", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, err
}

func (n *StringToLowerNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringToLowerNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringToLowerNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringToLowerNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		var converter block.NodeParameterConverter
		s, ok := converter.ToString(n.Data().InParameters.Get("input").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(strings.ToLower(s))
	}
	return value
}
