package branch

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	decimalRangeBranchNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "DecimalRangeBranchNode", FriendlyName: "Decimal Range Branch", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Base Condition"}}
	decimalRangeBranchNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Trigger a condition of a node execution based over a range of values (included in a min or max number)"}}
)

type DecimalRangeBranchNode struct {
	block.NodeBase
}

func NewDecimalRangeBranchNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(DecimalRangeBranchNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(value)

	rangeMin, err := block.NewNodeParameter(node, "rangeMin", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(rangeMin)

	rangeMax, err := block.NewNodeParameter(node, "rangeMax", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(rangeMax)

	inRange, err := block.NewNodeParameter(node, "In Range", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(inRange)

	lteMin, err := block.NewNodeParameter(node, "< RangeMin", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(lteMin)

	gteMax, err := block.NewNodeParameter(node, "> RangeMax", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(gteMax)

	return node, err
}

func (n *DecimalRangeBranchNode) CanBeExecuted() bool {
	return true
}

func (n *DecimalRangeBranchNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return decimalRangeBranchNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return decimalRangeBranchNodeGraphDescription
	default:
		return nil
	}
}

func (n *DecimalRangeBranchNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	value, ok := converter.ToDecimal(n.Data().InParameters.Get("value").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "value"}
	}

	rangeMin, ok := converter.ToDecimal(n.Data().InParameters.Get("rangeMin").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "rangeMin"}
	}

	rangeMax, ok := converter.ToDecimal(n.Data().InParameters.Get("rangeMax").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "rangeMax"}
	}

	if value.LessThan(rangeMin) {
		if outNode, ok := n.Data().OutParameters.Get("< RangeMin").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if value.GreaterThan(rangeMax) {
		if outNode, ok := n.Data().OutParameters.Get("> RangeMax").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	if value.GreaterThanOrEqual(rangeMin) && value.LessThanOrEqual(rangeMax) {
		if outNode, ok := n.Data().OutParameters.Get("In Range").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	return nil
}
