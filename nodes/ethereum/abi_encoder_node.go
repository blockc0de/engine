package ethereum

import (
	"context"
	"encoding/hex"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	abiEncoderNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AbiEncoderNode", FriendlyName: "Ethereum ABI Encoder", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	abiEncoderNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Encoding data params from ethereum transactions"}}
)

type AbiEncoderNode struct {
	block.NodeBase
}

func NewAbiEncoderNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(AbiEncoderNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	abi, err := block.NewNodeParameter(node, "abi", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(abi)

	method, err := block.NewNodeParameter(node, "method", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(method)

	params, err := block.NewNodeParameter(node, "params", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(params)

	data, err := block.NewDynamicNodeParameter(node, "data", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(data)

	return node, err
}

func (n *AbiEncoderNode) CanExecute() bool {
	return true
}

func (n *AbiEncoderNode) CanBeExecuted() bool {
	return true
}

func (n *AbiEncoderNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return abiEncoderNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return abiEncoderNodeGraphDescription
	default:
		return nil
	}
}

func (n *AbiEncoderNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("abi").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "abi"}
	}
	abiInstance, ok := value.(abi.ABI)
	if !ok {
		return block.ErrInvalidParameter{Name: "abi"}
	}

	var converter block.NodeParameterConverter
	method, ok := converter.ToString(n.Data().InParameters.Get("method").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "method"}
	}

	params, ok := converter.ToString(n.Data().InParameters.Get("params").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "params"}
	}

	encoder := NewAbiEncoder(abiInstance)
	data, err := encoder.Encode(method, params)
	if err != nil {
		return err
	}

	n.Data().OutParameters.Get("data").Value = block.NodeParameterString(hex.EncodeToString(data))

	return nil
}
