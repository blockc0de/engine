package ethereum

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethConnectionDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "EthConnection", FriendlyName: "Ethereum Connector", NodeType: attributes.NodeTypeEnumConnector, GroupName: "Blockchain.Ethereum", BlockLimitPerGraph: -1}}
	ethConnectionGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Connection to the Ethereum network, can be used as Managed connection (without in parameters) or with your own node"}}
)

type EthConnection struct {
	block.NodeBase
	Web3Client   *ethclient.Client `json:"-"`
	SocketClient *ethclient.Client `json:"-"`
}

func NewEthConnection(id string, graph *block.Graph) (block.Node, error) {
	node := new(EthConnection)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = false

	url, err := block.NewNodeParameter(node, "url", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(url)

	socketUrl, err := block.NewNodeParameter(node, "socketUrl", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(socketUrl)

	connection, err := block.NewDynamicNodeParameter(node, "connection", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(connection)

	return node, nil
}

func (n *EthConnection) CanExecute() bool {
	return true
}

func (n *EthConnection) SetupConnector(scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	url, ok := converter.ToString(n.Data().InParameters.Get("url").ComputeValue())
	if !ok {
		url = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	}

	socketUrl, ok := converter.ToString(n.Data().InParameters.Get("socketUrl").ComputeValue())
	if !ok {
		socketUrl = "wss://mainnet.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	}

	var err error
	n.Web3Client, err = ethclient.Dial(url)
	if err != nil {
		return err
	}

	n.SocketClient, err = ethclient.Dial(socketUrl)
	if err != nil {
		return err
	}

	scheduler.AddCycle(n, nil)
	return nil
}

func (n *EthConnection) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return ethConnectionDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return ethConnectionGraphDescription
	default:
		return nil
	}
}

func (n *EthConnection) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)
}

func (n *EthConnection) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("connection").Id {
		return n
	}
	return value
}

func (n *EthConnection) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *EthConnection) OnStop() error {
	return nil
}
