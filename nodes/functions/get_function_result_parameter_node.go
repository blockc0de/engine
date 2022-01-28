package functions

import (
	"context"
	"fmt"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	getFunctionResultParameterNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetFunctionResultParameterNode", FriendlyName: "Get Function Result Parameter", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	getFunctionResultParameterNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get a result parameter from a function result"}}
)

type GetFunctionResultParameterNode struct {
	block.NodeBase
}

func NewGetFunctionResultParameterNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetFunctionResultParameterNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	results, err := block.NewNodeParameter(node, "results", block.NodeParameterTypeEnumMapping, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(results)

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, nil
}

func (n *GetFunctionResultParameterNode) CanExecute() bool {
	return true
}

func (n *GetFunctionResultParameterNode) CanBeExecuted() bool {
	return true
}

func (n *GetFunctionResultParameterNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getFunctionResultParameterNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getFunctionResultParameterNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetFunctionResultParameterNode) OnExecution(context.Context, block.Engine) error {
	v := n.NodeData.InParameters.Get("results").ComputeValue()
	if v == nil {
		return block.ErrInvalidParameter{Name: "results"}
	}
	results, ok := v.(map[string]interface{})
	if !ok || results == nil {
		return block.ErrInvalidParameter{Name: "results"}
	}

	var converter block.NodeParameterConverter
	name, ok := converter.ToString(n.NodeData.InParameters.Get("name").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "name"}
	}

	value, ok := results[name]
	if !ok {
		return fmt.Errorf("no result named '%s' in function results", name)
	}

	n.NodeData.OutParameters.Get("value").Value = value

	return nil
}
