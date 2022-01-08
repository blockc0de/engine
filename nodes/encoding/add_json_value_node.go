package encoding

import (
	"context"
	"errors"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	addJsonValueNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddJsonValueNode", FriendlyName: "Add JSON Property", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON"}}
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
	var converter block.NodeParameterConverter
	key, ok := converter.ToString(n.NodeData.InParameters.Get("key").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "key"}
	}

	value, ok := converter.ToString(n.NodeData.InParameters.Get("value").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "value"}
	}

	v := n.NodeData.InParameters.Get("jsonObject").ComputeValue()
	if v == nil {
		return errors.New("invalid json object")
	}
	jsonObject, ok := v.(map[string]interface{})
	if !ok {
		return errors.New("invalid json object")
	}

	jsonObject[key] = value
	n.NodeData.OutParameters.Get("jsonObjectOut").Value = jsonObject
	return nil
}
