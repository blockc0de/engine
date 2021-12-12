package encoding

import (
	"context"
	"errors"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	jsoniter "github.com/json-iterator/go"
)

var (
	lastNodeToJsonNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "LastNodeToJsonNode", FriendlyName: "Last Node To JSON", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	lastNodeToJsonNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Serialize the last node executed in the graph to JSON data"}}
)

type LastNodeToJsonNode struct {
	block.NodeBase
}

func NewLastNodeToJsonNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(LastNodeToJsonNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	outParameter, err := block.NewDynamicNodeParameter(node, "json", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(outParameter)

	return node, nil
}

func (n *LastNodeToJsonNode) CanExecute() bool {
	return true
}

func (n *LastNodeToJsonNode) CanBeExecuted() bool {
	return true
}

func (n *LastNodeToJsonNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return lastNodeToJsonNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return lastNodeToJsonNodeGraphDescription
	default:
		return nil
	}
}

func (n *LastNodeToJsonNode) OnExecution(context.Context, block.NodeScheduler) error {
	if n.Data().LastExecutionFrom == nil {
		return errors.New("last node not found")
	}

	if !n.Data().LastExecutionFrom.Data().CanBeSerialized {
		n.NodeData.OutParameters.Get("json").Value = block.NodeParameterString(`{"error":"This node can't be serialized"}`)
	} else {
		object := make(map[string]interface{})
		for _, parameter := range n.Data().LastExecutionFrom.Data().OutParameters {
			object[parameter.Name] = parameter.ComputeValue()
		}

		data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&object)
		if err != nil {
			return err
		}
		n.NodeData.OutParameters.Get("json").Value = block.NodeParameterString(data)
	}

	return nil
}
