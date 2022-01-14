package storage

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/go-redis/redis"
)

var (
	delWalletKeyNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "DelWalletKeyNode", FriendlyName: "Delete Wallet Key Item", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Storage"}}
	delWalletKeyNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Delete a specific key in the Redis storage allocated for the wallet context."}}
)

type DelWalletKeyNode struct {
	block.NodeBase
	scope  string
	client redis.Cmdable
}

func NewDelWalletKeyNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(DelWalletKeyNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	key, err := block.NewNodeParameter(node, "key", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(key)

	return node, err
}

func (n *DelWalletKeyNode) CanExecute() bool {
	return true
}

func (n *DelWalletKeyNode) CanBeExecuted() bool {
	return true
}

func (n *DelWalletKeyNode) SetupDatabase(scope string, client redis.Cmdable) error {
	n.scope = scope
	n.client = client
	return nil
}

func (n *DelWalletKeyNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return delWalletKeyNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return delWalletKeyNodeGraphDescription
	default:
		return nil
	}
}

func (n *DelWalletKeyNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	key, ok := converter.ToString(n.Data().InParameters.Get("key").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "key"}
	}

	return n.client.HDel(redisWalletKey(n.scope), key).Err()
}
