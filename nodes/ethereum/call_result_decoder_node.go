package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	callResultDecoderNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CallResultDecoderNode", FriendlyName: "Call Result Decoder", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	callResultDecoderNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Decoding call result from ethereum transactions"}}
)

type CallResultDecoderNode struct {
	block.NodeBase
}

func NewCallResultDecoderNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CallResultDecoderNode)
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

	result, err := block.NewNodeParameter(node, "result", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(result)

	output, err := block.NewDynamicNodeParameter(node, "output", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(output)

	return node, err
}

func (n *CallResultDecoderNode) CanExecute() bool {
	return true
}

func (n *CallResultDecoderNode) CanBeExecuted() bool {
	return true
}

func (n *CallResultDecoderNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return callResultDecoderNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return callResultDecoderNodeGraphDescription
	default:
		return nil
	}
}

func (n *CallResultDecoderNode) OnExecution(ctx context.Context, engine block.Engine) error {
	value := n.Data().InParameters.Get("abi").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "abi"}
	}
	abiInstance, ok := value.(abi.ABI)
	if !ok {
		return block.ErrInvalidParameter{Name: "abi"}
	}

	var converter block.NodeParameterConverter
	result, ok := converter.ToString(n.Data().InParameters.Get("result").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "input"}
	}
	if strings.HasPrefix(result, "0x") || strings.HasPrefix(result, "0X") {
		result = result[2:]
	}

	method, ok := converter.ToString(n.Data().InParameters.Get("method").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "method"}
	}

	outputs, err := n.decodeMethodOutput(abiInstance, method, result)
	if err != nil {
		return err
	}

	data, err := json.Marshal(outputs)
	if err != nil {
		return err
	}

	n.Data().OutParameters.Get("output").Value = block.NodeParameterString(data)

	return nil
}

func (n *CallResultDecoderNode) decodeMethodOutput(abiInstance abi.ABI, methodName, result string) ([]interface{}, error) {
	data, err := hex.DecodeString(result)
	if err != nil {
		return nil, err
	}

	method, ok := abiInstance.Methods[methodName]
	if !ok {
		return nil, errors.New("method not found")
	}

	outputs, err := method.Outputs.Unpack(data)
	if err != nil {
		return nil, err
	}

	for idx, output := range outputs {
		if bn, ok := output.(*big.Int); ok {
			outputs[idx] = bn.String()
		}
	}
	return outputs, nil
}
