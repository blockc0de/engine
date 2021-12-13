package encoding

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	jsoniter "github.com/json-iterator/go"
)

var (
	mergeJsonNodeNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "MergeJsonNodeNode", FriendlyName: "Merge JSON", NodeType: attributes.NodeTypeEnumFunction, GroupName: "JSON", BlockLimitPerGraph: -1}}
	mergeJsonNodeNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Merge two JSON into one"}}
)

type MergeJsonNodeNode struct {
	block.NodeBase
}

func NewMergeJsonNodeNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(MergeJsonNodeNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	json1, err := block.NewNodeParameter(node, "json1", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(json1)

	json2, err := block.NewNodeParameter(node, "json2", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(json2)

	mergedJson, err := block.NewDynamicNodeParameter(node, "mergedJson", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(mergedJson)

	return node, nil
}

func (n *MergeJsonNodeNode) CanExecute() bool {
	return true
}

func (n *MergeJsonNodeNode) CanBeExecuted() bool {
	return true
}

func (n *MergeJsonNodeNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return mergeJsonNodeNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return mergeJsonNodeNodeGraphDescription
	default:
		return nil
	}
}

func (n *MergeJsonNodeNode) OnExecution(context.Context, block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	json1, ok := converter.ToString(n.NodeData.InParameters.Get("json1").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "json1"}
	}

	json2, ok := converter.ToString(n.NodeData.InParameters.Get("json2").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "json2"}
	}

	var object1 map[string]interface{}
	if err := json.Unmarshal([]byte(json1), &object1); err != nil {
		return err
	}

	var object2 map[string]interface{}
	if err := json.Unmarshal([]byte(json2), &object2); err != nil {
		return err
	}

	for key, val := range object2 {
		object1[key] = val
	}
	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(object1)
	if err != nil {
		return err
	}

	n.NodeData.OutParameters.Get("mergedJson").Value = block.NodeParameterString(data)

	return nil
}
