package time

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFormatTimestampNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	timestamp, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	timestamp.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(1639202246)}

	format, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	format.Data().OutParameters.Get("value").Value = block.NodeParameterString("%Y/%m/%d %H:%M:%S")

	formatTimestampNode, err := NewFormatTimestampNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	formatTimestampNode.Data().InParameters.Get("timestamp").Assignments = timestamp.Data().OutParameters.Get("value")
	formatTimestampNode.Data().InParameters.Get("format").Assignments = format.Data().OutParameters.Get("value")

	result := formatTimestampNode.ComputeParameterValue(formatTimestampNode.Data().OutParameters.Get("dateString").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "2021/12/11 05:57:26")
}
