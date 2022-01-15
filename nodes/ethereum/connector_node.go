package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethConnectionDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "EthConnection", FriendlyName: "Ethereum Connector", NodeType: attributes.NodeTypeEnumConnector, GroupName: "Blockchain.Ethereum"}}
	ethConnectionGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Connection to the Ethereum network, can be used as Managed connection (without in parameters) or with your own node"}}
)

type EthConnection struct {
	block.NodeBase
	ChainID          *big.Int          `json:"-"`
	IsSupportEIP1559 bool              `json:"-"`
	Web3Client       *ethclient.Client `json:"-"`
	SocketClient     *ethclient.Client `json:"-"`
}

func NewEthConnection(id string, graph *block.Graph) (block.Node, error) {
	node := new(EthConnection)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.CanBeSerialized = false

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
		url = os.Getenv("ETH_RPC_URL")
		if url == "" {
			url = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
		}
	}

	socketUrl, ok := converter.ToString(n.Data().InParameters.Get("socketUrl").ComputeValue())
	if !ok {
		socketUrl = os.Getenv("ETH_SOCKET_URL")
		if url == "" {
			url = "wss://mainnet.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161"
		}
	}

	var err error
	n.Web3Client, err = ethclient.Dial(url)
	if err != nil {
		return err
	}

	chainId1, err := n.getChainID(n.Web3Client)
	if err != nil {
		return err
	}

	n.SocketClient, err = ethclient.Dial(socketUrl)
	if err != nil {
		return err
	}

	chainId2, err := n.getChainID(n.SocketClient)
	if err != nil {
		return err
	}

	if chainId1.Cmp(chainId2) != 0 {
		return errors.New("chain id are not equal")
	}
	n.ChainID = chainId1

	n.IsSupportEIP1559, err = n.isSupportEIP1559(n.Web3Client)
	if err != nil {
		return err
	}

	scheduler.AppendLog("info",
		fmt.Sprintf("[%s] Successful connection to RPC node, chain: %d, eip1559: %v",
			n.Data().FriendlyName, chainId1.Int64(), n.IsSupportEIP1559))

	scheduler.AddCycle(n, nil)
	return nil
}

func (n *EthConnection) getChainID(client *ethclient.Client) (*big.Int, error) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	return client.ChainID(c)
}

func (n *EthConnection) isSupportEIP1559(client *ethclient.Client) (bool, error) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*30)
	blockNumber, err := client.BlockNumber(c)
	if err != nil {
		cancel()
		return false, err
	}
	cancel()

	c, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	block, err := client.BlockByNumber(c, big.NewInt(int64(blockNumber)))
	if err != nil {
		return false, err
	}
	return block.BaseFee() != nil, nil
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
	if n.Web3Client != nil {
		n.Web3Client.Close()
	}

	if n.SocketClient != nil {
		n.SocketClient.Close()
	}

	return nil
}
