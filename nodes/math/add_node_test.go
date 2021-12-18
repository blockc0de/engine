package math

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(10)}

	b, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(20)}

	add, err := NewAddNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	add.Data().InParameters.Get("a").Assignments = a.Data().OutParameters.Get("value")
	add.Data().InParameters.Get("b").Assignments = b.Data().OutParameters.Get("value")

	result := add.ComputeParameterValue(add.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "30")
}
