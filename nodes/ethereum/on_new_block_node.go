package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	onNewBlockEventNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnNewBlockEventNode", FriendlyName: "On Ethereum Block", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	onNewBlockEventNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Event that occurs everytime a new ethereum block is minted"}}
)

type OnNewBlockEventNode struct {
	block.NodeBase
	cancel       context.CancelFunc
	ch           chan *types.Header
	client       *ethclient.Client `json:"-"`
	subscription ethereum.Subscription
}

func NewOnNewBlockEventNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(OnNewBlockEventNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true

	connection, err := block.NewNodeParameter(node, "connection", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(connection)

	block, err := block.NewDynamicNodeParameter(node, "block", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(block)

	return node, nil
}

func (n *OnNewBlockEventNode) CanExecute() bool {
	return true
}

func (n *OnNewBlockEventNode) onNewBLock(scheduler block.NodeScheduler) {
	for {
		select {
		case header, ok := <-n.ch:
			if !ok {
				return
			}

			ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
			fullBlock, err := n.client.BlockByHash(ctx, header.Hash())
			if err != nil {
				scheduler.AppendLog("error", fmt.Sprintf("Failed to get block by hash, reason: %s", err.Error()))
			}

			v := struct {
				Header       *types.Header      `json:"header"`
				Transactions types.Transactions `json:"transactions"`
			}{header, fullBlock.Transactions()}
			data, err := json.Marshal(v)
			if err != nil {
				scheduler.AppendLog("error", fmt.Sprintf("Failed to marshal block, reason: %s", err.Error()))
			}

			p, err := block.NewDynamicNodeParameter(n, "block", block.NodeParameterTypeEnumString, false)
			if err != nil {
				scheduler.AppendLog("error", fmt.Sprintf("Failed to create dynamic node parameter, reason: %s", err.Error()))
			}
			p.Value = block.NodeParameterString(data)
			scheduler.AddCycle(n, []*block.NodeParameter{p})
		}
	}
}

func (n *OnNewBlockEventNode) SetupEvent(scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("connection").ComputeValue()
	if value == nil {
		return errors.New("invalid connection")
	}
	connection, ok := value.(*EthConnection)
	if !ok {
		return errors.New("invalid connection")
	}

	var err error
	n.ch = make(chan *types.Header, 64)
	n.client = connection.SocketClient
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	n.subscription, err = connection.SocketClient.SubscribeNewHead(ctx, n.ch)
	if err != nil {
		return err
	}

	go n.onNewBLock(scheduler)

	return nil
}

func (n *OnNewBlockEventNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return onNewBlockEventNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return onNewBlockEventNodeGraphDescription
	default:
		return nil
	}
}

func (n *OnNewBlockEventNode) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)
}

func (n *OnNewBlockEventNode) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *OnNewBlockEventNode) OnStop() error {
	if n.subscription != nil {
		n.subscription.Unsubscribe()
		close(n.ch)
	}
	return nil
}
