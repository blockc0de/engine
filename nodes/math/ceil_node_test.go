package math

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCeilNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	number, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	number.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(10.123)}

	ceil, err := NewCeilNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	ceil.Data().InParameters.Get("number").Assignments = number.Data().OutParameters.Get("value")

	result := ceil.ComputeParameterValue(ceil.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "11")
}
