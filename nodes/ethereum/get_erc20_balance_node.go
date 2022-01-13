package ethereum

import (
	"context"
	"errors"
	"math/big"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/ethereum/web3util"
	"github.com/ethereum/go-ethereum/common"
)

var (
	getErc20BalanceNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetERC20BalanceNode", FriendlyName: "Get ERC20 Balance", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	getErc20BalanceNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get the balance of an address for a ERC20"}}
)

type GetErc20BalanceNode struct {
	block.NodeBase
}

func NewGetErc20BalanceNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetErc20BalanceNode)
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

	address, err := block.NewNodeParameter(node, "address", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(address)

	balance, err := block.NewDynamicNodeParameter(node, "balance", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(balance)

	return node, nil
}

func (n *GetErc20BalanceNode) CanExecute() bool {
	return true
}

func (n *GetErc20BalanceNode) CanBeExecuted() bool {
	return true
}

func (n *GetErc20BalanceNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getErc20BalanceNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getErc20BalanceNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetErc20BalanceNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
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

	address, ok := converter.ToString(n.Data().InParameters.Get("address").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "address"}
	}

	decimals, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "decimals")
	if err != nil {
		return err
	}
	d, ok := decimals[0].(uint8)
	if !ok {
		return errors.New("invalid decimals")
	}

	balance, err := erc20Call(ctx, connection.Web3Client, common.HexToAddress(contract), "balanceOf", common.HexToAddress(address))
	if err != nil {
		return err
	}
	bal, ok := balance[0].(*big.Int)
	if !ok {
		return errors.New("invalid balance")
	}

	n.NodeData.OutParameters.Get("balance").Value = block.NodeParameterDecimal{Decimal: web3util.WeiToDecimal(bal, int(d))}

	return nil
}
