package web3util

import (
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

var (
	hexToIntegerDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "HexToIntegerNode", FriendlyName: "Hex To Integer", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Web3.Util", BlockLimitPerGraph: -1}}
	hexToIntegerGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Returns the number representation of a given HEX value."}}
)

type HexToIntegerNode struct {
	block.NodeBase
}

func NewHexToIntegerNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(HexToIntegerNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	hex, err := block.NewNodeParameter(node, "hex", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(hex)

	integer, err := block.NewDynamicNodeParameter(node, "integer", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(integer)

	return node, err
}

func (n *HexToIntegerNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return hexToIntegerDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return hexToIntegerGraphDescription
	default:
		return nil
	}
}

func (n *HexToIntegerNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("integer").Id {
		hex := n.Data().InParameters.Get("hex")
		var converter block.NodeParameterConverter
		hexVal, ok := converter.ToString(hex.ComputeValue())
		if !ok {
			return nil
		}

		hexVal = strings.ToLower(hexVal)
		hexVal = strings.TrimPrefix(hexVal, "0x")

		for len(hexVal) > 0 && hexVal[0] == '0' {
			hexVal = hexVal[1:]
		}

		bn, err := hexutil.DecodeBig("0x" + hexVal)
		if err != nil {
			return nil
		}

		return block.NodeParameterDecimal{Decimal: decimal.NewFromBigInt(bn, 0)}
	}
	return value
}
