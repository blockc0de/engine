package ethereum

import (
	"bytes"
	"context"
	"math/big"
	"reflect"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var (
	erc20TokenAbi                         *abi.ABI
	erc20TokenJsonAbi                     = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
	getErc20TokenInfoNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetERC20TokenInformationsNode", FriendlyName: "Get ERC20 Informations", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	getErc20TokenInfoNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get informations about a ERC20 token"}}
)

func init() {
	a, err := abi.JSON(bytes.NewReader([]byte(erc20TokenJsonAbi)))
	if err != nil {
		panic(err)
	}
	erc20TokenAbi = &a
}

type GetErc20TokenInfoNode struct {
	block.NodeBase
}

func NewGetErc20TokenInfoNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetErc20TokenInfoNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true

	connection, err := block.NewNodeParameter(node, "connection", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(connection)

	contract, err := block.NewNodeParameter(node, "contract", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(contract)

	name, err := block.NewDynamicNodeParameter(node, "name", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(name)

	symbol, err := block.NewDynamicNodeParameter(node, "symbol", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(symbol)

	decimals, err := block.NewDynamicNodeParameter(node, "decimals", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(decimals)

	totalSupply, err := block.NewDynamicNodeParameter(node, "totalSupply", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(totalSupply)

	return node, nil
}

func (n *GetErc20TokenInfoNode) CanExecute() bool {
	return true
}

func (n *GetErc20TokenInfoNode) CanBeExecuted() bool {
	return true
}

func (n *GetErc20TokenInfoNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getErc20TokenInfoNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getErc20TokenInfoNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetErc20TokenInfoNode) ethCall(ctx context.Context, client *ethclient.Client,
	contract common.Address, name string, args ...interface{}) ([]interface{}, error) {

	data, err := erc20TokenAbi.Pack(name, args...)
	if err != nil {
		return nil, err
	}

	c, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	msg := ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}
	output, err := client.CallContract(c, msg, nil)
	if err != nil {
		return nil, err
	}

	return erc20TokenAbi.Unpack(name, output)
}

func (n *GetErc20TokenInfoNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("connection").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "connection"}
	}
	connection, ok := value.(*EthConnection)
	if !ok {
		return block.ErrInvalidParameter{Name: "connection"}
	}

	var converter block.NodeParameterConverter
	contract, ok := converter.ToString(n.Data().InParameters.Get("contract").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "contract"}
	}

	name, err := n.ethCall(ctx, connection.Web3Client, common.HexToAddress(contract), "name")
	if err != nil {
		return err
	}
	if s, ok := name[0].(string); ok {
		n.NodeData.OutParameters.Get("name").Value = block.NodeParameterString(s)
	}

	symbol, err := n.ethCall(ctx, connection.Web3Client, common.HexToAddress(contract), "symbol")
	if err != nil {
		return err
	}
	if s, ok := symbol[0].(string); ok {
		n.NodeData.OutParameters.Get("symbol").Value = block.NodeParameterString(s)
	}

	decimals, err := n.ethCall(ctx, connection.Web3Client, common.HexToAddress(contract), "decimals")
	if err != nil {
		return err
	}
	if d, ok := decimals[0].(uint8); ok {
		n.NodeData.OutParameters.Get("decimals").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(int64(d))}
	}

	totalSupply, err := n.ethCall(ctx, connection.Web3Client, common.HexToAddress(contract), "totalSupply")
	if err != nil {
		return err
	}
	if d, ok := totalSupply[0].(*big.Int); ok {
		n.NodeData.OutParameters.Get("totalSupply").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromBigInt(d, 0)}
	}

	return nil
}
