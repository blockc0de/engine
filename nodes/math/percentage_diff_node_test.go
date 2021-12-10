package math

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/variable"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPercentageDiffNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(10)}

	b, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(12)}

	percentageDiffNode, err := NewPercentageDiffNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	percentageDiffNode.Data().InParameters.Get("a").Assignments = a.Data().OutParameters.Get("value")
	percentageDiffNode.Data().InParameters.Get("b").Assignments = b.Data().OutParameters.Get("value")

	result := percentageDiffNode.ComputeParameterValue(percentageDiffNode.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "20")
}
