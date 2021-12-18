package encoding

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	jsoniter "github.com/json-iterator/go"
)

var (
	convertToJsonNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "ConvertToJsonNode", FriendlyName: "Convert To JSON", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	convertToJsonNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Convert the received any type parameter into a JSON object readable"}}
)

type ConvertToJsonNode struct {
	block.NodeBase
}

func NewConvertToJsonNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(ConvertToJsonNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	object, err := block.NewNodeParameter(node, "object", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(object)

	json, err := block.NewDynamicNodeParameter(node, "json", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(json)

	return node, nil
}

func (n *ConvertToJsonNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return convertToJsonNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return convertToJsonNodeGraphDescription
	default:
		return nil
	}
}

func (n *ConvertToJsonNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("json").Id {
		v := n.Data().InParameters.Get("object").ComputeValue()
		if v == nil {
			return nil
		}

		data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
		if err != nil {
			return nil
		}
		return block.NodeParameterString(data)
	}
	return value
}
