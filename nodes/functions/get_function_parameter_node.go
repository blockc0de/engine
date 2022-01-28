package functions

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	getFunctionParameterNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetFunctionParameterNode", FriendlyName: "Get Function Parameter", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Function"}}
	getFunctionParameterNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Get a call parameter from the current function context"}}
)

type GetFunctionParameterNode struct {
	block.NodeBase
}

func NewGetFunctionParameterNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetFunctionParameterNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	name, err := block.NewNodeParameter(node, "name", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(name)

	value, err := block.NewDynamicNodeParameter(node, "value", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(value)

	return node, nil
}

func (n *GetFunctionParameterNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getFunctionParameterNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getFunctionParameterNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetFunctionParameterNode) ComputeParameterValue(id string, value interface{}) interface{} {
	if id == n.NodeData.OutParameters.Get("value").Id {
		var converter block.NodeParameterConverter
		name, ok := converter.ToString(n.NodeData.InParameters.Get("name").ComputeValue())
		if !ok {
			return block.ErrInvalidParameter{Name: "name"}
		}

		val := n.NodeData.Graph.CurrentCycle.LocalStorage.Get(CurrentFunctionContext)
		context, ok := val.(*FunctionContext)
		if !ok || context == nil || context.CallParameters == nil {
			return nil
		}

		parameter, ok := context.CallParameters[name]
		if !ok {
			return nil
		}
		return parameter
	}
	return value
}
