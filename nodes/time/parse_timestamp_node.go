package time

import (
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/itchyny/timefmt-go"
	"github.com/shopspring/decimal"
)

var (
	parseTimestampNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "ParseTimestamp", FriendlyName: "Parse Timestamp", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Time", BlockLimitPerGraph: -1}}
	parseTimestampNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Parse a timestamp with a given pattern"}}
)

type ParseTimestampNode struct {
	block.NodeBase
}

func NewParseTimestampNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(ParseTimestampNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	dateString, err := block.NewNodeParameter(node, "dateString", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(dateString)

	format, err := block.NewNodeParameter(node, "format", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(format)

	timestamp, err := block.NewDynamicNodeParameter(node, "timestamp", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(timestamp)

	return node, err
}

func (n *ParseTimestampNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return parseTimestampNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return parseTimestampNodeGraphDescription
	default:
		return nil
	}
}

func (n *ParseTimestampNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("timestamp").Id {
		var converter block.NodeParameterConverter
		dateString, ok := converter.ToString(n.Data().InParameters.Get("dateString").ComputeValue())
		if !ok {
			return nil
		}

		format, ok := converter.ToString(n.Data().InParameters.Get("format").ComputeValue())
		if !ok {
			return nil
		}

		t, err := timefmt.Parse(dateString, format)
		if err != nil {
			return nil
		}
		return block.NodeParameterDecimal{Decimal: decimal.NewFromInt(t.Unix())}
	}
	return value
}
