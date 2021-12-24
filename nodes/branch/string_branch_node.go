package branch

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringBranchNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringBranchNode", FriendlyName: "String Branch", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Base Condition", BlockLimitPerGraph: -1}}
	stringBranchNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Trigger a condition over two string value to compare if their are equals to each other"}}
)

type StringBranchNode struct {
	block.NodeBase
}

func NewStringBranchNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringBranchNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	valueA, err := block.NewNodeParameter(node, "valueA", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(valueA)

	valueB, err := block.NewNodeParameter(node, "valueB", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(valueB)

	eq, err := block.NewNodeParameter(node, "==", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(eq)

	notEq, err := block.NewNodeParameter(node, "!=", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(notEq)

	return node, err
}

func (n *StringBranchNode) CanBeExecuted() bool {
	return true
}

func (n *StringBranchNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringBranchNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringBranchNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringBranchNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	valueA, ok := converter.ToString(n.Data().InParameters.Get("valueA").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "valueA"}
	}

	valueB, ok := converter.ToString(n.Data().InParameters.Get("valueB").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "valueB"}
	}

	if valueA == valueB {
		if outNode, ok := n.Data().OutParameters.Get("==").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, scheduler)
		}
	} else {
		if outNode, ok := n.Data().OutParameters.Get("!=").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, scheduler)
		}
	}

	return nil
}
