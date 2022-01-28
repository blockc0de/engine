package functions

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	addFunctionResultNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddFunctionResultNode", FriendlyName: "Set Function Result", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	addFunctionResultNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Set a value returned by the function"}}
)

type AddFunctionResultNode struct {
	block.NodeBase
}

func NewAddFunctionResultNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AddFunctionResultNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(value)

	return node, nil
}

func (n *AddFunctionResultNode) CanExecute() bool {
	return true
}

func (n *AddFunctionResultNode) CanBeExecuted() bool {
	return true
}

func (n *AddFunctionResultNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return addFunctionResultNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return addFunctionResultNodeGraphDescription
	default:
		return nil
	}
}

func (n *AddFunctionResultNode) OnExecution(context.Context, block.Engine) error {
	var converter block.NodeParameterConverter
	name, ok := converter.ToString(n.NodeData.InParameters.Get("name").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "name"}
	}

	val := n.NodeData.Graph.CurrentCycle.LocalStorage.Get(CurrentFunctionContext)
	context, ok := val.(*FunctionContext)
	if !ok || context == nil || context.ReturnValues == nil {
		return nil
	}

	context.ReturnValues[name] = n.NodeData.InParameters.Get("value").ComputeValue()

	return nil
}
