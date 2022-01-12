package ethereum

import (
	"context"
	"math/big"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var (
	getErc20TokenInfoNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetERC20TokenInformationsNode", FriendlyName: "Get ERC20 Informations", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	getErc20TokenInfoNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get informations about a ERC20 token"}}
)

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

	name, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "name")
	if err != nil {
		return err
	}
	if s, ok := name[0].(string); ok {
		n.NodeData.OutParameters.Get("name").Value = block.NodeParameterString(s)
	}

	symbol, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "symbol")
	if err != nil {
		return err
	}
	if s, ok := symbol[0].(string); ok {
		n.NodeData.OutParameters.Get("symbol").Value = block.NodeParameterString(s)
	}

	decimals, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "decimals")
	if err != nil {
		return err
	}
	if d, ok := decimals[0].(uint8); ok {
		n.NodeData.OutParameters.Get("decimals").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(int64(d))}
	}

	totalSupply, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "totalSupply")
	if err != nil {
		return err
	}
	if d, ok := totalSupply[0].(*big.Int); ok {
		n.NodeData.OutParameters.Get("totalSupply").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromBigInt(d, 0)}
	}

	return nil
}
