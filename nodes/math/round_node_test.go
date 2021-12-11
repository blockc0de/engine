package math

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	number, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	number.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(5.55)}

	places, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	places.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(1)}

	round, err := NewRoundNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	round.Data().InParameters.Get("number").Assignments = number.Data().OutParameters.Get("value")
	round.Data().InParameters.Get("places").Assignments = places.Data().OutParameters.Get("value")

	result := round.ComputeParameterValue(round.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "5.6")
}
