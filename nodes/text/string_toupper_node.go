package text

import (
	"github.com/graphlinq-go/engine/attributes"
	"github.com/graphlinq-go/engine/block"
	"reflect"
	"strings"
)

var (
	stringToUpperNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringToUpperNode", FriendlyName: "String To Upper", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String", BlockLimitPerGraph: -1}}
	stringToUpperNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns s with all Unicode letters mapped to their upper case."}}
)

type StringToUpperNode struct {
	block.NodeBase
}

func NewStringToUpperNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringToUpperNode)
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

func (n *StringToUpperNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringToUpperNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringToUpperNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringToUpperNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		var converter block.NodeParameterConverter
		str := n.Data().InParameters.Get("input")
		stringVal, ok := converter.ToString(str.ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(strings.ToUpper(stringVal))
	}
	return value
}
