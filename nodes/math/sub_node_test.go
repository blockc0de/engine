package math

import (
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(99)}

	b, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(9)}

	sub, err := NewSubNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	sub.Data().InParameters.Get("a").Assignments = a.Data().OutParameters.Get("value")
	sub.Data().InParameters.Get("b").Assignments = b.Data().OutParameters.Get("value")

	result := sub.ComputeParameterValue(sub.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "90")
}
