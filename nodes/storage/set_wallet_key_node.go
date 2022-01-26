package storage

import (
	"context"
	"fmt"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/go-redis/redis"
)

var (
	setWalletKeyNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "SetWalletKeyNode", FriendlyName: "Set Wallet Key Item", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Storage"}}
	setWalletKeyNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Save a specific key in the Redis storage allocated for the wallet context."}}
)

type SetWalletKeyNode struct {
	block.NodeBase
	scope  string
	client redis.Cmdable
}

func NewSetWalletKeyNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(SetWalletKeyNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	key, err := block.NewNodeParameter(node, "key", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(key)

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(value)

	return node, err
}

func (n *SetWalletKeyNode) CanExecute() bool {
	return true
}

func (n *SetWalletKeyNode) CanBeExecuted() bool {
	return true
}

func (n *SetWalletKeyNode) SetupDatabase(scope string, client redis.Cmdable) error {
	n.scope = scope
	n.client = client
	return nil
}

func (n *SetWalletKeyNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return setWalletKeyNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return setWalletKeyNodeGraphDescription
	default:
		return nil
	}
}

func (n *SetWalletKeyNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	key, ok := converter.ToString(n.Data().InParameters.Get("key").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "key"}
	}

	value, ok := converter.ToString(n.Data().InParameters.Get("value").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "value"}
	}

	count, err := n.client.HLen(redisWalletKey(n.scope)).Result()
	if err != nil {
		return err
	}

	if count >= KeysLimitPerWallet {
		return fmt.Errorf("keys limit per wallet, want: %d", KeysLimitPerWallet)
	}
	return n.client.HSet(redisWalletKey(n.scope), key, value).Err()
}
