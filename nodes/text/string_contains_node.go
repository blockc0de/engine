package text

import (
	"context"
	"errors"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"reflect"
	"strings"
)

var (
	stringContainsNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "StringContainsNode", FriendlyName: "String Contains", NodeType: attributes.NodeTypeEnumCondition, GroupName: "String", BlockLimitPerGraph: -1}}
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

	retTrue, err := block.NewNodeParameter(node, "true", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(retTrue)

	retFalse, err := block.NewNodeParameter(node, "false", block.NodeParameterTypeEnumNode, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(retFalse)

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

func (n *StringContainsNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	inParameterString := n.Data().InParameters.Get("string")
	inParameterToSearch := n.Data().InParameters.Get("toSearch")

	var converter block.NodeParameterConverter
	stringVal, ok := converter.ToString(inParameterString.ComputeValue())
	if !ok {
		return errors.New("invalid parameter")
	}

	toSearchVal, ok := converter.ToString(inParameterToSearch.ComputeValue())
	if !ok {
		return errors.New("invalid parameter")
	}

	if strings.Index(strings.ToLower(stringVal), strings.ToLower(toSearchVal)) >= 0 {
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
