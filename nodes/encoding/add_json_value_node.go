package encoding

import (
	"context"
	"errors"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	addJsonValueNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddJsonValueNode", FriendlyName: "Add JSON Property", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	addJsonValueNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Add a property into a JSON Object"}}
)

type AddJsonValueNode struct {
	block.NodeBase
}

func NewAddJsonValueNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AddJsonValueNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	jsonObject, err := block.NewNodeParameter(node, "jsonObject", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(jsonObject)

	key, err := block.NewNodeParameter(node, "key", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(key)

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(value)

	jsonObjectOut, err := block.NewDynamicNodeParameter(node, "jsonObjectOut", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(jsonObjectOut)

	return node, nil
}

func (n *AddJsonValueNode) CanExecute() bool {
	return true
}

func (n *AddJsonValueNode) CanBeExecuted() bool {
	return true
}

func (n *AddJsonValueNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return addJsonValueNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return addJsonValueNodeGraphDescription
	default:
		return nil
	}
}

func (n *AddJsonValueNode) OnExecution(context.Context, block.NodeScheduler) error {
	key := n.NodeData.InParameters.Get("key")
	value := n.NodeData.InParameters.Get("value")

	var converter block.NodeParameterConverter
	keyVal, ok := converter.ToString(key.ComputeValue())
	if !ok {
		return errors.New("invalid parameter")
	}

	valueVal, ok := converter.ToString(value.ComputeValue())
	if !ok {
		return errors.New("invalid parameter")
	}

	v := n.NodeData.InParameters.Get("jsonObject").ComputeValue()
	if v == nil {
		return errors.New("invalid json object")
	}
	jsonObject, ok := v.(map[string]interface{})
	if !ok {
		return errors.New("invalid json object")
	}

	jsonObject[keyVal] = valueVal
	n.NodeData.OutParameters.Get("jsonObjectOut").Value = jsonObject
	return nil
}
