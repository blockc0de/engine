package encoding

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/tidwall/gjson"
)

var (
	jsonSelectorNodeNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "JsonSelectorNodeNode", FriendlyName: "JSON Selector", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON"}}
	jsonSelectorNodeNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Select a specific value in a json object and return it as string parameter"}}
)

type JsonSelectorNodeNode struct {
	block.NodeBase
}

func NewJsonSelectorNodeNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(JsonSelectorNodeNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	json, err := block.NewNodeParameter(node, "json", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(json)

	selector, err := block.NewNodeParameter(node, "selector", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(selector)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, nil
}

func (n *JsonSelectorNodeNode) CanExecute() bool {
	return true
}

func (n *JsonSelectorNodeNode) CanBeExecuted() bool {
	return true
}

func (n *JsonSelectorNodeNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return jsonSelectorNodeNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return jsonSelectorNodeNodeGraphDescription
	default:
		return nil
	}
}

func (n *JsonSelectorNodeNode) OnExecution(context.Context, block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	js, ok := converter.ToString(n.NodeData.InParameters.Get("json").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "json"}
	}

	selector, ok := converter.ToString(n.NodeData.InParameters.Get("selector").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "selector"}
	}

	value := gjson.Get(js, selector)
	if value.Type == gjson.Null {
		n.NodeData.OutParameters.Get("value").Value = block.NodeParameterString("null")
	} else {
		n.NodeData.OutParameters.Get("value").Value = block.NodeParameterString(value.String())
	}

	return nil
}
