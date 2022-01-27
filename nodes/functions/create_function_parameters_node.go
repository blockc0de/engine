package functions

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

type FunctionParameters map[string]interface{}

var (
	createFunctionParametersNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CreateFunctionParametersNode", FriendlyName: "Create Function Parameters", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	createFunctionParametersNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Create a empty array for function parameters"}}
)

type CreateFunctionParametersNode struct {
	block.NodeBase
}

func NewCreateFunctionParametersNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CreateFunctionParametersNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	parameters, err := block.NewDynamicNodeParameter(node, "parameters", block.NodeParameterTypeEnumMapping, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(parameters)

	return node, nil
}

func (n *CreateFunctionParametersNode) CanExecute() bool {
	return true
}

func (n *CreateFunctionParametersNode) CanBeExecuted() bool {
	return true
}

func (n *CreateFunctionParametersNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return createFunctionParametersNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return createFunctionParametersNodeGraphDescription
	default:
		return nil
	}
}

func (n *CreateFunctionParametersNode) OnExecution(context.Context, block.Engine) error {
	parameters := make(FunctionParameters)
	n.NodeData.OutParameters.Get("parameters").Value = parameters
	return nil
}
