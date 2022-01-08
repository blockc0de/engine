package ethereum

import (
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var ERC20ABI abi.ABI
var ERC20AbiJSON = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

func init() {
	var err error
	ERC20ABI, err = abi.JSON(strings.NewReader(ERC20AbiJSON))
	if err != nil {
		panic(err)
	}
}

var (
	erc20AbiNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "Erc20AbiNode", FriendlyName: "ERC20 ABI", NodeType: attributes.NodeTypeEnumVariable, GroupName: "Blockchain.Ethereum"}}
	erc20AbiNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "ABI for ERC20 contract on Ethereum"}}
)

type Erc20AbiNode struct {
	block.NodeBase
}

func NewErc20AbiNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(Erc20AbiNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	abi, err := block.NewDynamicNodeParameter(node, "abi", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(abi)

	return node, err
}

func (n *Erc20AbiNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return erc20AbiNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return erc20AbiNodeGraphDescription
	default:
		return nil
	}
}

func (n *Erc20AbiNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("abi").Id {
		return ERC20ABI
	}
	return value
}
