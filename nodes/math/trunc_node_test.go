package math

import (
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	number, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	number.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(5.5566)}

	precision, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	precision.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(3)}

	trunc, err := NewTruncNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	trunc.Data().InParameters.Get("number").Assignments = number.Data().OutParameters.Get("value")
	trunc.Data().InParameters.Get("precision").Assignments = precision.Data().OutParameters.Get("value")

	result := trunc.ComputeParameterValue(trunc.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "5.556")
}
