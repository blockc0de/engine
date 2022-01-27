package encoding

import (
	"encoding/json"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	jsonDeserializeToArrayNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "JsonDeserializeToArrayNode", FriendlyName: "JSON Deserialize To Array", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON"}}
	jsonDeserializeToArrayNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Convert a plain json string to a array"}}
)

type JsonDeserializeToArrayNode struct {
	block.NodeBase
}

func NewJsonDeserializeToArrayNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(JsonDeserializeToArrayNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	json, err := block.NewNodeParameter(node, "json", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(json)

	array, err := block.NewDynamicNodeParameter(node, "array", block.NodeParameterTypeEnumArray, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(array)

	return node, nil
}

func (n *JsonDeserializeToArrayNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return jsonDeserializeToArrayNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return jsonDeserializeToArrayNodeGraphDescription
	default:
		return nil
	}
}

func (n *JsonDeserializeToArrayNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("array").Id {
		var converter block.NodeParameterConverter
		data, ok := converter.ToString(n.Data().InParameters.Get("json").ComputeValue())
		if !ok {
			return nil
		}

		var array []interface{}
		if err := json.Unmarshal([]byte(data), &array); err != nil {
			return nil
		}
		return &array
	}
	return value
}
