package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
)

var (
	sendTransactionNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "SendTransactionNode", FriendlyName: "Send Transaction", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Blockchain.Ethereum"}}
	sendTransactionNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Sends a transaction to the network."}}
)

type SendTransactionNode struct {
	block.NodeBase
}

func NewSendTransactionNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(SendTransactionNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	connection, err := block.NewNodeParameter(node, "connection", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(connection)

	key, err := block.NewNodeParameter(node, "key", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(key)

	to, err := block.NewNodeParameter(node, "to", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(to)

	data, err := block.NewNodeParameter(node, "data", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(data)

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(value)

	nonce, err := block.NewNodeParameter(node, "nonce", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(nonce)

	gasPrice, err := block.NewNodeParameter(node, "gasPrice", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(gasPrice)

	gasLimit, err := block.NewNodeParameter(node, "gasLimit", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(gasLimit)

	hash, err := block.NewDynamicNodeParameter(node, "hash", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(hash)

	return node, nil
}

func (n *SendTransactionNode) CanExecute() bool {
	return true
}

func (n *SendTransactionNode) CanBeExecuted() bool {
	return true
}

func (n *SendTransactionNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return sendTransactionNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return sendTransactionNodeGraphDescription
	default:
		return nil
	}
}

func (n *SendTransactionNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	conn := n.Data().InParameters.Get("connection").ComputeValue()
	if conn == nil {
		return block.ErrInvalidParameter{Name: "connection"}
	}
	connection, ok := conn.(*EthConnection)
	if !ok {
		return block.ErrInvalidParameter{Name: "connection"}
	}

	var converter block.NodeParameterConverter
	key, ok := converter.ToString(n.Data().InParameters.Get("key").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "key"}
	}

	privateKey, fromAddress, err := n.parsePrivateKey(key)
	if err != nil {
		return err
	}

	var toAddress *common.Address
	to, ok := converter.ToString(n.Data().InParameters.Get("to").ComputeValue())
	if ok {
		address := common.HexToAddress(to)
		toAddress = &address
	}

	var hexData []byte
	data, ok := converter.ToString(n.Data().InParameters.Get("data").ComputeValue())
	if ok {
		if strings.HasPrefix(data, "0x") || strings.HasPrefix(data, "0X") {
			data = data[2:]
		}
		if hexData, err = hex.DecodeString(data); err != nil {
			return err
		}
	}

	value, _ := converter.ToDecimal(n.Data().InParameters.Get("value").ComputeValue())
	nonce, _ := converter.ToDecimal(n.Data().InParameters.Get("nonce").ComputeValue())
	gasPrice, _ := converter.ToDecimal(n.Data().InParameters.Get("gasPrice").ComputeValue())
	gasLimit, _ := converter.ToDecimal(n.Data().InParameters.Get("gasLimit").ComputeValue())

	// Get Account nonce
	if nonce.IsZero() {
		pending := big.NewInt(-1)
		n, err := connection.Web3Client.NonceAt(ctx, fromAddress, pending)
		if err != nil {
			return err
		}
		nonce = decimal.NewFromBigInt(big.NewInt(0).SetUint64(n), 0)
	}

	// Estimate gas
	if gasLimit.IsZero() {
		msg := ethereum.CallMsg{
			To:   toAddress,
			Data: hexData,
		}
		gas, err := connection.Web3Client.EstimateGas(ctx, msg)
		if err != nil {
			return err
		}
		gasLimit = decimal.NewFromBigInt(big.NewInt(0).SetUint64(gas), 0)
	}

	// Get gas price
	if gasPrice.IsZero() {
		bn, err := connection.Web3Client.SuggestGasPrice(ctx)
		if err != nil {
			return err
		}
		gasPrice = decimal.NewFromBigInt(bn, 0)
	}

	// Create transaction
	var tx *types.Transaction
	if !connection.IsSupportEIP1559 {
		legacyTx := types.LegacyTx{
			To:       toAddress,
			Gas:      gasLimit.BigInt().Uint64(),
			GasPrice: gasPrice.BigInt(),
			Value:    value.BigInt(),
			Data:     hexData,
			Nonce:    nonce.BigInt().Uint64(),
		}
		tx = types.NewTx(&legacyTx)
	} else {
		dynamicFeeTx := types.DynamicFeeTx{
			ChainID:   connection.ChainID,
			To:        toAddress,
			Gas:       gasLimit.BigInt().Uint64(),
			GasTipCap: big.NewInt(params.GWei),
			GasFeeCap: gasPrice.BigInt(),
			Value:     value.BigInt(),
			Data:      hexData,
			Nonce:     nonce.BigInt().Uint64(),
		}
		tx = types.NewTx(&dynamicFeeTx)
	}

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(connection.ChainID), privateKey)
	if err != nil {
		return err
	}

	// Send signed transaction
	err = connection.Web3Client.SendTransaction(ctx, signedTx)
	if err != nil {
		scheduler.AppendLog("error", fmt.Sprintf(
			"[%s] Failed to send transaction, from: %s, to: %s, hash: %s, reason: %s",
			n.Data().FriendlyName, fromAddress.String(), toAddress.String(), tx.Hash().String(), err.Error()))
		return err
	}

	n.NodeData.OutParameters.Get("hash").Value = block.NodeParameterString(tx.Hash().String())

	scheduler.AppendLog("info", fmt.Sprintf(
		"[%s] Broadcast transaction, from: %s, to: %s, hash: %s",
		n.Data().FriendlyName, fromAddress.String(), toAddress.String(), tx.Hash().String()))

	return nil
}

func (n *SendTransactionNode) parsePrivateKey(key string) (*ecdsa.PrivateKey, common.Address, error) {
	if strings.HasPrefix(key, "0x") || strings.HasPrefix(key, "0X") {
		key = key[2:]
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
