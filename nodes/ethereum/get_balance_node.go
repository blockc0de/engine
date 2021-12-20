package ethereum

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/ethereum/web3util"
	"github.com/ethereum/go-ethereum/common"
)

var (
	getBalanceNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetEtherBalanceNode", FriendlyName: "Get Ether Balance", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	getBalanceNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get the balance of address in ether"}}
)

type GetBalanceNode struct {
	block.NodeBase
}

func NewGetBalanceNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetBalanceNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true

	connection, err := block.NewNodeParameter(node, "connection", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(connection)

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

func (n *GetBalanceNode) CanExecute() bool {
	return true
}

func (n *GetBalanceNode) CanBeExecuted() bool {
	return true
}

func (n *GetBalanceNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getBalanceNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getBalanceNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetBalanceNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("connection").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "connection"}
	}
	connection, ok := value.(*EthConnection)
	if !ok {
		return block.ErrInvalidParameter{Name: "connection"}
	}

	var converter block.NodeParameterConverter
	address, ok := converter.ToString(n.Data().InParameters.Get("address").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "address"}
	}

	balance, err := connection.Web3Client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return err
	}

	n.NodeData.OutParameters.Get("balance").Value = block.NodeParameterDecimal{
		Decimal: web3util.WeiToDecimal(balance, 18),
	}

	return nil
}
