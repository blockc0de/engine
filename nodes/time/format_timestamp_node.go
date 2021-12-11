package time

import (
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"github.com/itchyny/timefmt-go"
	"reflect"
	"time"
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
		timestamp := n.Data().InParameters.Get("timestamp")
		format := n.Data().InParameters.Get("format")

		var converter block.NodeParameterConverter
		timestampVal, ok := converter.ToDecimal(timestamp.ComputeValue())
		if !ok {
			return nil
		}

		formatVal, ok := converter.ToString(format.ComputeValue())
		if !ok {
			return nil
		}

		str := timefmt.Format(time.Unix(timestampVal.IntPart(), 0).UTC(), formatVal)
		return block.NodeParameterString(str)
	}
	return value
}
