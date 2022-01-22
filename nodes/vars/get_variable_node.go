package vars

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	getVariableNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetVariable", FriendlyName: "Get variable", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Base Variable"}}
	getVariableNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Return the value of the variable pre computed from a Set variable block"}}
)

type GetVariableNode struct {
	block.NodeBase
}

func NewGetVariableNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetVariableNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.CanBeSerialized = false

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	value, err := block.NewNodeParameter(node, "value", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, err
}

func (n *GetVariableNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getVariableNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getVariableNodeGraphDescription
	default:
		return nil
	}
}
func (n *GetVariableNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("value").Id {
		name := n.Data().InParameters.Get("name")

		var converter block.NodeParameterConverter
		nameVal, ok := converter.ToString(name.ComputeValue())
		if !ok {
			return nil
		}

		value, ok := n.Data().Graph.MemoryVariables[nameVal]
		if !ok {
			return nil
		}
		return value
	}
	return value
}
