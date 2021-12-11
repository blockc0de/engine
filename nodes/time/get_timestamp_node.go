package time

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

var (
	getTimestampNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "GetTimestampNode", FriendlyName: "Get Timestamp", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Time", BlockLimitPerGraph: -1}}
	getTimestampNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Return the current timestamp of the engine localtime"}}
)

type GetTimestampNode struct {
	block.NodeBase
}

func NewGetTimestampNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(GetTimestampNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	timestamp, err := block.NewDynamicNodeParameter(node, "timestamp", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(timestamp)

	return node, err
}

func (n *GetTimestampNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return getTimestampNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return getTimestampNodeGraphDescription
	default:
		return nil
	}
}

func (n *GetTimestampNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("timestamp").Id {
		return block.NodeParameterDecimal{Decimal: decimal.NewFromInt(time.Now().Unix())}
	}
	return value
}
