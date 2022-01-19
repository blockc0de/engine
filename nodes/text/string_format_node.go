package text

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringFormatNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringFormatNode", FriendlyName: "Format String", NodeType: attributes.NodeTypeEnumFunction, GroupName: "String"}}
	stringFormatNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Formats according to a format specifier and returns the resulting string."}}
)

type StringFormatNode struct {
	block.NodeBase
}

func NewStringFormatNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringFormatNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	format, err := block.NewNodeParameter(node, "format", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(format)

	args, err := block.NewNodeParameter(node, "args", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(args)

	outParameter, err := block.NewDynamicNodeParameter(node, "string", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, err
}

func (n *StringFormatNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringFormatNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringFormatNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringFormatNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("string").Id {
		var converter block.NodeParameterConverter
		format, ok := converter.ToString(n.Data().InParameters.Get("format").ComputeValue())
		if !ok {
			return nil
		}

		args := n.Data().InParameters.Get("args").ComputeValue()
		if s, ok := converter.ToString(args); ok && json.Valid([]byte(s)) {
			var err error
			x := strings.TrimLeft(s, " \t\r\n")
			isArray := len(x) > 0 && x[0] == '['
			if isArray {
				args = make([]interface{}, 0)
				err = json.Unmarshal([]byte(s), &args)
			}

			isObject := len(x) > 0 && x[0] == '{'
			if isObject {
				args = make(map[string]interface{})
				err = json.Unmarshal([]byte(s), &args)
			}

			if err != nil {
				return nil
			}
		}
		return block.NodeParameterString(StringFormatter{}.Format(format, args))
	}
	return value
}
