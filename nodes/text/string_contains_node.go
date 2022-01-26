package text

import (
	"context"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	stringContainsNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringContainsNode", FriendlyName: "String Contains", NodeType: attributes.NodeTypeEnumCondition, GroupName: "String"}}
	stringContainsNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Check if a string contains a other string"}}
)

type StringContainsNode struct {
	block.NodeBase
}

func NewStringContainsNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(StringContainsNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	str, err := block.NewNodeParameter(node, "string", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(str)

	toSearch, err := block.NewNodeParameter(node, "toSearch", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(toSearch)

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

func (n *StringContainsNode) CanBeExecuted() bool {
	return true
}

func (n *StringContainsNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return stringContainsNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return stringContainsNodeGraphDescription
	default:
		return nil
	}
}

func (n *StringContainsNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	s, ok := converter.ToString(n.Data().InParameters.Get("string").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "string"}
	}

	toSearch, ok := converter.ToString(n.Data().InParameters.Get("toSearch").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "toSearch"}
	}

	if strings.Contains(strings.ToLower(s), strings.ToLower(toSearch)) {
		if outNode, ok := n.Data().OutParameters.Get("true").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	} else {
		if outNode, ok := n.Data().OutParameters.Get("false").Value.(block.ExecutableNode); ok && outNode != nil {
			return outNode.OnExecution(ctx, engine)
		}
	}

	return nil
}
