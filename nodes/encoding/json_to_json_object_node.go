package encoding

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	jsoniter "github.com/json-iterator/go"
)

var (
	jsonToJsonObjectNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "JsonToJsonObjectNode", FriendlyName: "JSON to JSON Object", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	jsonToJsonObjectNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Convert a plain json string to a json object"}}
)

type JsonToJsonObjectNode struct {
	block.NodeBase
}

func NewJsonToJsonObjectNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(JsonToJsonObjectNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	json, err := block.NewNodeParameter(node, "json", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(json)

	jsonObject, err := block.NewDynamicNodeParameter(node, "jsonObject", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(jsonObject)

	return node, nil
}

func (n *JsonToJsonObjectNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return jsonToJsonObjectNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return jsonToJsonObjectNodeGraphDescription
	default:
		return nil
	}
}

func (n *JsonToJsonObjectNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("jsonObject").Id {
		var converter block.NodeParameterConverter
		data, ok := converter.ToString(n.Data().InParameters.Get("json").ComputeValue())
		if !ok {
			return nil
		}

		object := make(map[string]interface{})
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		if err := json.Unmarshal([]byte(data), &object); err != nil {
			return nil
		}

		return object
	}
	return value
}
