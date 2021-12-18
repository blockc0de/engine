package vars

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	setVariableNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "SetVariable", FriendlyName: "Set variable", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Base Variable", BlockLimitPerGraph: -1}}
	setVariableNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Set a variable value (can be any type) in the graph memory context"}}
)

type SetVariableNode struct {
	block.NodeBase
}

func NewSetVariableNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(SetVariableNode)
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

	return node, err
}

func (n *SetVariableNode) CanExecute() bool {
	return true
}

func (n *SetVariableNode) CanBeExecuted() bool {
	return true
}

func (n *SetVariableNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return setVariableNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return setVariableNodeGraphDescription
	default:
		return nil
	}
}
func (n *SetVariableNode) OnExecution(context.Context, block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	name := n.Data().InParameters.Get("name")
	nameVal, ok := converter.ToString(name.ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "name"}
	}

	value := n.Data().InParameters.Get("value")
	n.Data().Graph.MemoryVariables[nameVal] = value.ComputeValue()
	return nil
}
