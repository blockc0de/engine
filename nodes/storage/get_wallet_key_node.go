package storage

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/go-redis/redis"
)

var (
	getWalletKeyNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetWalletKeyNode", FriendlyName: "Get Wallet Key Item", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Storage"}}
	getWalletKeyNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Return a specific key from the Redis storage allocated for the wallet context."}}
)

type GetWalletKeyNode struct {
	block.NodeBase
	scope  string
	client redis.Cmdable
}

func NewGetWalletKeyNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetWalletKeyNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	key, err := block.NewNodeParameter(node, "key", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(key)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *GetWalletKeyNode) CanExecute() bool {
	return true
}

func (n *GetWalletKeyNode) CanBeExecuted() bool {
	return true
}

func (n *GetWalletKeyNode) SetupDatabase(scope string, client redis.Cmdable) error {
	n.scope = scope
	n.client = client
	return nil
}

func (n *GetWalletKeyNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getWalletKeyNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getWalletKeyNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetWalletKeyNode) OnExecution(ctx context.Context, engine block.Engine) error {
	var converter block.NodeParameterConverter
	key, ok := converter.ToString(n.Data().InParameters.Get("key").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "key"}
	}

	value, err := n.client.HGet(redisWalletKey(n.scope), key).Result()
	if err != nil {
		return err
	}

	n.NodeData.OutParameters.Get("value").Value = block.NodeParameterString(value)

	return nil
}
