package vars

import (
	"context"
	"errors"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
)

var (
	isVariableExistNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "IsVariableExist", FriendlyName: "Is Variable Exist", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Base Variable", BlockLimitPerGraph: -1}}
	isVariableExistNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Check if a variable exist in the graph memory context"}}
)

type IsVariableExistNode struct {
	block.NodeBase
}

func NewIsVariableExistNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(IsVariableExistNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	t, err := block.NewNodeParameter(node, "true", block.NodeParameterTypeEnumNode, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(t)

	f, err := block.NewNodeParameter(node, "false", block.NodeParameterTypeEnumNode, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(f)

	return node, err
}

func (n *IsVariableExistNode) CanBeExecuted() bool {
	return true
}

func (n *IsVariableExistNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return isVariableExistNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return isVariableExistNodeGraphDescription
	default:
		return nil
	}
}
func (n *IsVariableExistNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	name := n.Data().InParameters.Get("name")
	nameVal, ok := converter.ToString(name.ComputeValue())
	if !ok {
		return errors.New("invalid parameter")
	}

	_, exist := n.Data().Graph.MemoryVariables[nameVal]
	if exist {
		if outNode, ok := n.Data().OutParameters.Get("true").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, scheduler)
		}
	} else {
		if outNode, ok := n.Data().OutParameters.Get("false").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, scheduler)
		}
	}

	return nil
}
