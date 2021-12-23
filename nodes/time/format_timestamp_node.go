package time

import (
	"reflect"
	"time"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/itchyny/timefmt-go"
)

var (
	formatTimestampNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "FormatTimestamp", FriendlyName: "Format Timestamp", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Time", BlockLimitPerGraph: -1}}
	formatTimestampNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Format a timestamp with a given pattern"}}
)

type FormatTimestampNode struct {
	block.NodeBase
}

func NewFormatTimestampNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(FormatTimestampNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	timestamp, err := block.NewNodeParameter(node, "timestamp", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(timestamp)

	format, err := block.NewNodeParameter(node, "format", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(format)

	dateString, err := block.NewDynamicNodeParameter(node, "dateString", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(dateString)

	return node, err
}

func (n *FormatTimestampNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return formatTimestampNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return formatTimestampNodeGraphDescription
	default:
		return nil
	}
}

func (n *FormatTimestampNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("dateString").Id {
		var converter block.NodeParameterConverter
		timestamp, ok := converter.ToDecimal(n.Data().InParameters.Get("timestamp").ComputeValue())
		if !ok {
			return nil
		}

		format, ok := converter.ToString(n.Data().InParameters.Get("format").ComputeValue())
		if !ok {
			return nil
		}

		str := timefmt.Format(time.Unix(timestamp.IntPart(), 0).UTC(), format)
		return block.NodeParameterString(str)
	}
	return value
}
