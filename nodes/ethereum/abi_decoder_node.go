package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	abiDecoderNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "AbiDecoderNode", FriendlyName: "Ethereum ABI Decoder", NodeType: attributes.NodeTypeEnumCondition, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	abiDecoderNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Decoding data params from ethereum transactions"}}
)

type methodInput struct {
	Method string        `json:"method"`
	Inputs []interface{} `json:"inputs"`
}

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

	result, err := block.NewNodeParameter(node, "result", block.NodeParameterTypeEnumString, false, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(result)

	return node, err
}

func (n *AbiDecoderNode) CanExecute() bool {
	return true
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
	value := n.Data().InParameters.Get("abi").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "abi"}
	}
	abiInstance, ok := value.(abi.ABI)
	if !ok {
		return block.ErrInvalidParameter{Name: "abi"}
	}

	var converter block.NodeParameterConverter
	input, ok := converter.ToString(n.Data().InParameters.Get("input").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "input"}
	}
	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		input = input[2:]
	}
	if len(input) < 8 {
		return block.ErrInvalidParameter{Name: "input"}
	}

	result, err := n.decodeMethodInput(abiInstance, input)
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	n.Data().OutParameters.Get("result").Value = block.NodeParameterString(data)

	return nil
}

func (n *AbiDecoderNode) decodeMethodInput(abiInstance abi.ABI, input string) (methodInput, error) {
	data, err := hex.DecodeString(input)
	if err != nil {
		return methodInput{}, err
	}

	method, err := abiInstance.MethodById(data)
	if err != nil {
		return methodInput{}, err
	}

	inputs, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return methodInput{}, err
	}

	for idx, input := range inputs {
		if bn, ok := input.(*big.Int); ok {
			inputs[idx] = bn.String()
		}
	}
	return methodInput{Method: method.RawName, Inputs: inputs}, nil
}
