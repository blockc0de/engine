package ethereum

import (
	"context"
	"encoding/hex"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var (
	callContractNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CallContractNode", FriendlyName: "Call Contract", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	callContractNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Executes a new message call immediately without creating a transaction on the block chain."}}
)

type CallContractNode struct {
	block.NodeBase
}

func NewCallContractNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CallContractNode)
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

	data, err := block.NewNodeParameter(node, "data", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(data)

	output, err := block.NewDynamicNodeParameter(node, "output", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(output)

	return node, nil
}

func (n *CallContractNode) CanExecute() bool {
	return true
}

func (n *CallContractNode) CanBeExecuted() bool {
	return true
}

func (n *CallContractNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return callContractNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return callContractNodeGraphDescription
	default:
		return nil
	}
}

func (n *CallContractNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
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

	data, ok := converter.ToString(n.Data().InParameters.Get("data").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "data"}
	}
	if strings.HasPrefix(data, "0x") || strings.HasPrefix(data, "0X") {
		data = data[2:]
	}

	hexData, err := hex.DecodeString(data)
	if err != nil {
		return err
	}

	to := common.HexToAddress(contract)
	msg := ethereum.CallMsg{
		To:   &to,
		Data: hexData,
	}
	output, err := connection.Web3Client.CallContract(ctx, msg, nil)
	if err != nil {
		return err
	}

	n.NodeData.OutParameters.Get("output").Value = block.NodeParameterString(hex.EncodeToString(output))

	return nil
}
