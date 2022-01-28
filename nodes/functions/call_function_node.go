package functions

import (
	"context"
	"fmt"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	callFunctionNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CallFunctionNode", FriendlyName: "Call Function", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	callFunctionNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Call a function in the graph"}}
)

type CallFunctionNode struct {
	block.NodeBase
}

func NewCallFunctionNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CallFunctionNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	parameters, err := block.NewNodeParameter(node, "parameters", block.NodeParameterTypeEnumMapping, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(parameters)

	results, err := block.NewDynamicNodeParameter(node, "results", block.NodeParameterTypeEnumMapping, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(results)

	return node, nil
}

func (n *CallFunctionNode) CanExecute() bool {
	return true
}

func (n *CallFunctionNode) CanBeExecuted() bool {
	return true
}

func (n *CallFunctionNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return callFunctionNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return callFunctionNodeGraphDescription
	default:
		return nil
	}
}

func (n *CallFunctionNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	name, ok := converter.ToString(n.NodeData.InParameters.Get("name").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "name"}
	}

	v := n.NodeData.InParameters.Get("parameters").ComputeValue()
	if v == nil {
		return block.ErrInvalidParameter{Name: "parameters"}
	}
	parameters, ok := v.(FunctionParameters)
	if !ok || parameters == nil {
		return block.ErrInvalidParameter{Name: "parameters"}
	}

	functionNode := n.Data().Graph.FindNode(func(node block.Node) bool {
		if node.Data().NodeType != functionNodeType {
			return false
		}

		n, ok := converter.ToString(node.Data().InParameters.Get("name").ComputeValue())
		if !ok {
			return false
		}

		return n == name
	})
	if functionNode == nil {
		return fmt.Errorf("function '%s' doesnt exist in the graph", name)
	}

	executableNode := functionNode.(*FunctionNode)
	executableNode.CallParameters = parameters
	if err := executableNode.OnExecution(ctx, engine); err != nil {
		return err
	}

	n.NodeData.OutParameters.Get("results").Value = executableNode.Context.ReturnValues

	return nil
}
