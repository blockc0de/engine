package encoding

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	createJsonObjectNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "CreateJsonObjectNode", FriendlyName: "Create JSON Object", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	createJsonObjectNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Create a new empty JSON object"}}
)

type CreateJsonObjectNode struct {
	block.NodeBase
}

func NewCreateJsonObjectNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(CreateJsonObjectNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	jsonObject, err := block.NewDynamicNodeParameter(node, "jsonObject", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(jsonObject)

	return node, nil
}

func (n *CreateJsonObjectNode) CanExecute() bool {
	return true
}

func (n *CreateJsonObjectNode) CanBeExecuted() bool {
	return true
}

func (n *CreateJsonObjectNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return createJsonObjectNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return createJsonObjectNodeGraphDescription
	default:
		return nil
	}
}

func (n *CreateJsonObjectNode) OnExecution(context.Context, block.NodeScheduler) error {
	jsonObject := make(map[string]interface{}, 0)
	n.NodeData.OutParameters.Get("jsonObject").Value = jsonObject
	return nil
}
