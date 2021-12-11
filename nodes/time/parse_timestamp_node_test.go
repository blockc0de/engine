package time

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTimestampNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	dateString, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	dateString.Data().OutParameters.Get("value").Value = block.NodeParameterString("2021/12/11 13:57:26")

	format, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	format.Data().OutParameters.Get("value").Value = block.NodeParameterString("%Y/%m/%d %H:%M:%S")

	parseTimestampNode, err := NewParseTimestampNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	parseTimestampNode.Data().InParameters.Get("dateString").Assignments = dateString.Data().OutParameters.Get("value")
	parseTimestampNode.Data().InParameters.Get("format").Assignments = format.Data().OutParameters.Get("value")

	result := parseTimestampNode.ComputeParameterValue(parseTimestampNode.Data().OutParameters.Get("timestamp").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).IntPart(), int64(1639231046))
}
