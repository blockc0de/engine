package time

import (
	"math"
	"testing"
	"time"

	"github.com/blockc0de/engine/block"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTimestamp(t *testing.T) {
	graph := block.NewGraph("", "test")

	getTimestampNode, err := NewGetTimestampNode(uuid.New().String(), graph)
	assert.Nil(t, err)

	result := getTimestampNode.ComputeParameterValue(getTimestampNode.Data().OutParameters.Get("timestamp").Id, nil)
	assert.True(t, math.Abs(float64(result.(block.NodeParameterDecimal).IntPart()-time.Now().Unix())) < 10)
}
