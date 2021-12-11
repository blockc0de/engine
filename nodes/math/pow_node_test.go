package math

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(10)}

	b, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(0.5)}

	pow, err := NewPowNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	pow.Data().InParameters.Get("a").Assignments = a.Data().OutParameters.Get("value")
	pow.Data().InParameters.Get("b").Assignments = b.Data().OutParameters.Get("value")

	result := pow.ComputeParameterValue(pow.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "3.16227766018")
}
