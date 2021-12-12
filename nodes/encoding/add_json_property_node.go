package encoding

import (
	"context"
	"errors"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	addJsonPropertyNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddJsonPropertyNode", FriendlyName: "Add JSON Property", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	addJsonPropertyNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Add a property into a JSON Object"}}
)

type AddJsonPropertyNode struct {
	block.NodeBase
}

func NewAddJsonPropertyNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AddJsonPropertyNode)
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

func (n *AddJsonPropertyNode) CanExecute() bool {
	return true
}

func (n *AddJsonPropertyNode) CanBeExecuted() bool {
	return true
}

func (n *AddJsonPropertyNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return addJsonPropertyNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return addJsonPropertyNodeGraphDescription
	default:
		return nil
	}
}

func (n *AddJsonPropertyNode) OnExecution(context.Context, block.NodeScheduler) error {
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

	jsonObject, ok := n.NodeData.InParameters.Get("jsonObject").Value.(map[string]interface{})
	if !ok {
		return errors.New("invalid json object")
	}

	jsonObject[keyVal] = valueVal
	return nil
}
