package branch

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	decimalBranchNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "DecimalBranchNode", FriendlyName: "Decimal Branch", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Base Condition"}}
	decimalBranchNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Trigger different node path on a condition based on a decimal variable value state (equals to, greater or equals, lower then, lower or equals..)."}}
)

type DecimalBranchNode struct {
	block.NodeBase
}

func NewDecimalBranchNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(DecimalBranchNode)
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

	gt, err := block.NewNodeParameter(node, ">", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(gt)

	gte, err := block.NewNodeParameter(node, ">=", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(gte)

	eq, err := block.NewNodeParameter(node, "==", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(eq)

	lte, err := block.NewNodeParameter(node, "<=", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(lte)

	lt, err := block.NewNodeParameter(node, "<", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(lt)

	return node, err
}

func (n *DecimalBranchNode) CanBeExecuted() bool {
	return true
}

func (n *DecimalBranchNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return decimalBranchNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return decimalBranchNodeGraphDescription
	default:
		return nil
	}
}

func (n *DecimalBranchNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	valueA, ok := converter.ToDecimal(n.Data().InParameters.Get("valueA").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "valueA"}
	}

	valueB, ok := converter.ToDecimal(n.Data().InParameters.Get("valueB").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "valueB"}
	}

	if valueA.GreaterThan(valueB) {
		if outNode, ok := n.Data().OutParameters.Get(">").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if valueA.GreaterThanOrEqual(valueB) {
		if outNode, ok := n.Data().OutParameters.Get(">=").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if valueA.Equal(valueB) {
		if outNode, ok := n.Data().OutParameters.Get("==").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if valueA.LessThanOrEqual(valueB) {
		if outNode, ok := n.Data().OutParameters.Get("<=").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if valueA.LessThan(valueB) {
		if outNode, ok := n.Data().OutParameters.Get("<").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	return nil
}
