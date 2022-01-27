package functions

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	adaddFunctionParameterNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AddFunctionParameterNode", FriendlyName: "Add Function Parameter", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	adaddFunctionParameterNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Add a new parameter to a function parameters array"}}
)

type AdaddFunctionParameterNode struct {
	block.NodeBase
}

func NewAdaddFunctionParameterNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AdaddFunctionParameterNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	parameters, err := block.NewNodeParameter(node, "parameters", block.NodeParameterTypeEnumMapping, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(parameters)

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

	outParameters, err := block.NewDynamicNodeParameter(node, "outParameters", block.NodeParameterTypeEnumMapping, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameters)

	return node, nil
}

func (n *AdaddFunctionParameterNode) CanExecute() bool {
	return true
}

func (n *AdaddFunctionParameterNode) CanBeExecuted() bool {
	return true
}

func (n *AdaddFunctionParameterNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return adaddFunctionParameterNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return adaddFunctionParameterNodeGraphDescription
	default:
		return nil
	}
}

func (n *AdaddFunctionParameterNode) OnExecution(context.Context, block.Engine) error {
	v := n.NodeData.InParameters.Get("parameters").ComputeValue()
	if v == nil {
		return block.ErrInvalidParameter{Name: "parameters"}
	}
	parameters, ok := v.(FunctionParameters)
	if !ok || parameters == nil {
		return block.ErrInvalidParameter{Name: "parameters"}
	}

	var converter block.NodeParameterConverter
	name, ok := converter.ToString(n.NodeData.InParameters.Get("name").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "name"}
	}

	value := n.NodeData.InParameters.Get("value").ComputeValue()

	parameters[name] = value
	n.NodeData.OutParameters.Get("outParameters").Value = parameters

	return nil
}
