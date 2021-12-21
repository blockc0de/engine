package ethereum

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	abiDecoderNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AbiDecoderNode", FriendlyName: "Ethereum ABI Decoder", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	abiDecoderNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Decoding data params from ethereum transactions"}}
)

type AbiDecoderNode struct {
	block.NodeBase
}

func NewAbiDecoderNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AbiDecoderNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	input, err := block.NewNodeParameter(node, "input", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(input)

	abi, err := block.NewNodeParameter(node, "abi", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(abi)

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

	result, err := block.NewNodeParameter(node, "result", block.NodeParameterTypeEnumString, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(result)

	return node, err
}

func (n *AbiDecoderNode) CanBeExecuted() bool {
	return true
}

func (n *AbiDecoderNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return abiDecoderNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return abiDecoderNodeGraphDescription
	default:
		return nil
	}
}

func (n *AbiDecoderNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	return nil
}
