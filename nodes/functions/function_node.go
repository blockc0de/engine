package functions

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

const (
	CurrentFunctionContext = "CurrentFunctionContext"
)

var (
	functionNodeType             = reflect.TypeOf(new(FunctionNode)).String()
	functionNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "FunctionNode", FriendlyName: "Function", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	functionNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Create a new function"}}
)

type FunctionNode struct {
	block.NodeBase
	Context        *FunctionContext
	CallParameters FunctionParameters
}

func NewFunctionNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(FunctionNode)
	node.NodeData = block.NewNodeData(id, node, graph, functionNodeType)

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	return node, err
}

func (n *FunctionNode) CanExecute() bool {
	return true
}

func (n *FunctionNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return functionNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return functionNodeGraphDescription
	default:
		return nil
	}
}

func (n *FunctionNode) OnExecution(ctx context.Context, engine block.Engine) error {
	n.Context = NewFunctionContext(n)
	n.Context.CallParameters = n.CallParameters

	n.NodeData.Graph.CurrentCycle.LocalStorage.Add(CurrentFunctionContext, n.Context)

	engine.NextNode(ctx, n)

	return nil
}
