package ethereum

import (
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	jsonAbiNodeDefinition             = []interface{}{attributes.NodeDefinition{NodeName: "JsonAbiNode", FriendlyName: "JSON ABI", NodeType: attributes.NodeTypeEnumVariable, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	jsonAbiNodeGraphDescription       = []interface{}{attributes.NodeGraphDescription{Description: "A string that contain a JSON ABI"}}
	jsonAbiNodeIDEParametersAttribute = []interface{}{attributes.NodeIDEParametersAttribute{IsScriptInput: true, ScriptType: "json"}}
)

type JsonAbiNode struct {
	block.NodeBase
	oldValue    interface{}
	contractAbi abi.ABI
}

func NewJsonAbiNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(JsonAbiNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	abi, err := block.NewNodeParameter(node, "abi", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(abi)

	return node, err
}

func (n *JsonAbiNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return jsonAbiNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return jsonAbiNodeGraphDescription
	case reflect.TypeOf(attributes.NodeIDEParametersAttribute{}):
		return jsonAbiNodeIDEParametersAttribute
	default:
		return nil
	}
}

func (n *JsonAbiNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("abi").Id {
		if value == nil {
			return nil
		}

		if n.oldValue == value {
			return n.contractAbi
		}

		var converter block.NodeParameterConverter
		abiVal, ok := converter.ToString(value)
		if !ok {
			return nil
		}

		contractAbi, err := abi.JSON(strings.NewReader(abiVal))
		if err != nil {
			return nil
		}

		n.oldValue = value
		n.contractAbi = contractAbi

		return contractAbi
	}
	return value
}
