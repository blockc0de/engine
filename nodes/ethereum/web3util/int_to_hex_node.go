package web3util

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	integerToHexDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "IntegerToHexNode", FriendlyName: "Integer To Hex", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Web3.Util"}}
	integerToHexGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns the HEX representation of a given number value."}}
)

type IntegerToHexNode struct {
	block.NodeBase
}

func NewIntegerToHexNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(IntegerToHexNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	integer, err := block.NewNodeParameter(node, "integer", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(integer)

	hex, err := block.NewDynamicNodeParameter(node, "hex", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(hex)

	return node, err
}

func (n *IntegerToHexNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return integerToHexDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return integerToHexGraphDescription
	default:
		return nil
	}
}

func (n *IntegerToHexNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("hex").Id {
		var converter block.NodeParameterConverter
		integer, ok := converter.ToDecimal(n.Data().InParameters.Get("integer").ComputeValue())
		if !ok {
			return nil
		}

		return block.NodeParameterString(hexutil.EncodeBig(integer.BigInt()))
	}
	return value
}
