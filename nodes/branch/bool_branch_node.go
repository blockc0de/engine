package branch

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	boolBranchNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "BoolBranchNode", FriendlyName: "Bool Branch", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Base Condition", BlockLimitPerGraph: -1}}
	boolBranchNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Trigger different node path on a condition based on a bool variable value state (true/false)."}}
)

type BoolBranchNode struct {
	block.NodeBase
}

func NewBoolBranchNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(BoolBranchNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	condition, err := block.NewNodeParameter(node, "condition", block.NodeParameterTypeEnumBool, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(condition)

	t, err := block.NewNodeParameter(node, "true", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(t)

	f, err := block.NewNodeParameter(node, "false", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(f)

	return node, err
}

func (n *BoolBranchNode) CanBeExecuted() bool {
	return true
}

func (n *BoolBranchNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return boolBranchNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return boolBranchNodeGraphDescription
	default:
		return nil
	}
}

func (n *BoolBranchNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	condition, ok := converter.ToBool(n.Data().InParameters.Get("condition").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "condition"}
	}

	if condition {
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
