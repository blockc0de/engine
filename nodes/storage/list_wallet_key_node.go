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
	listWalletKeyNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "ListWalletKeyNode", FriendlyName: "List Wallet Key Item", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Storage"}}
	listWalletKeyNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Return all keys from the Redis storage allocated for the wallet context."}}
)

type ListWalletKeyNode struct {
	block.NodeBase
	scope  string
	client redis.Cmdable
}

func NewListWalletKeyNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(ListWalletKeyNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	array, err := block.NewDynamicNodeParameter(node, "array", block.NodeParameterTypeEnumArray, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(array)

	return node, err
}

func (n *ListWalletKeyNode) CanExecute() bool {
	return true
}

func (n *ListWalletKeyNode) CanBeExecuted() bool {
	return true
}

func (n *ListWalletKeyNode) SetupDatabase(scope string, client redis.Cmdable) error {
	n.scope = scope
	n.client = client
	return nil
}

func (n *ListWalletKeyNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return listWalletKeyNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return listWalletKeyNodeGraphDescription
	default:
		return nil
	}
}

func (n *ListWalletKeyNode) OnExecution(ctx context.Context, engine block.Engine) error {
	mapper, err := n.client.HGetAll(redisWalletKey(n.scope)).Result()
	if err != nil {
		return err
	}

	array := make([]interface{}, 0, len(mapper))
	format := `{"key":"%s","value":"%s"}`
	for key, value := range mapper {
		array = append(array, fmt.Sprintf(format, key, value))
	}

	n.NodeData.OutParameters.Get("array").Value = array

	return nil
}
