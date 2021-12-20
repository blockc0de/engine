package ethereum

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	onEventLogNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnEventLogNode", FriendlyName: "On Ethereum Event Log", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	onEventLogNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Event that occurs everytime a new contract event log is emitted"}}
)

type OnEventLogNode struct {
	block.NodeBase
	ch           chan types.Log
	client       *ethclient.Client
	subscription ethereum.Subscription
}

func NewOnEventLogNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(OnEventLogNode)
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

	eventLog, err := block.NewDynamicNodeParameter(node, "eventLog", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(eventLog)

	return node, nil
}

func (n *OnEventLogNode) CanExecute() bool {
	return true
}

func (n *OnEventLogNode) handleRead(scheduler block.NodeScheduler) {
	for {
		select {
		case eventLog, ok := <-n.ch:
			if !ok {
				return
			}

			data, err := eventLog.MarshalJSON()
			if err != nil {
				scheduler.AppendLog("error", fmt.Sprintf("Failed to marshal event log, reason: %s", err.Error()))
				break
			}

			p, err := block.NewDynamicNodeParameter(n, "eventLog", block.NodeParameterTypeEnumString, false)
			if err != nil {
				scheduler.AppendLog("error", fmt.Sprintf("Failed to create dynamic node parameter, reason: %s", err.Error()))
				break
			}
			p.Value = block.NodeParameterString(data)
			scheduler.AddCycle(n, []*block.NodeParameter{p})
		}
	}
}

func (n *OnEventLogNode) SetupEvent(scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("connection").ComputeValue()
	if value == nil {
		return errors.New("invalid connection")
	}
	connection, ok := value.(*EthConnection)
	if !ok {
		return errors.New("invalid connection")
	}

	var converter block.NodeParameterConverter
	contract, ok := converter.ToString(n.Data().InParameters.Get("contract").ComputeValue())
	if !ok {
		return errors.New("invalid contract address")
	}

	var err error
	n.ch = make(chan types.Log, 64)
	n.client = connection.SocketClient

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	filter := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(contract)}}
	n.subscription, err = connection.SocketClient.SubscribeFilterLogs(ctx, filter, n.ch)
	if err != nil {
		return err
	}

	go n.handleRead(scheduler)

	return nil
}

func (n *OnEventLogNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return onEventLogNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return onEventLogNodeGraphDescription
	default:
		return nil
	}
}

func (n *OnEventLogNode) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)
}

func (n *OnEventLogNode) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *OnEventLogNode) OnStop() error {
	if n.subscription != nil {
		n.subscription.Unsubscribe()
		close(n.ch)
	}
	return nil
}
