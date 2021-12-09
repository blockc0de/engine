package math

import (
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloorNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	number, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	number.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(2.3323546)}

	floor, err := NewFloorNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	floor.Data().InParameters.Get("number").Assignments = number.Data().OutParameters.Get("value")

	result := floor.ComputeParameterValue(floor.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "2")
}
