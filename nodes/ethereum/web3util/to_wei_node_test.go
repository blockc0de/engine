package web3util

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDecimalToWei(t *testing.T) {
	graph := block.NewGraph("", "test")

	value, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	value.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromFloat(1.024)}

	decimals, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	decimals.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(18)}

	toWeiNode, err := NewToWeiNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	toWeiNode.Data().InParameters.Get("value").Assignments = value.Data().OutParameters.Get("value")
	toWeiNode.Data().InParameters.Get("decimals").Assignments = decimals.Data().OutParameters.Get("value")

	result := toWeiNode.ComputeParameterValue(toWeiNode.Data().OutParameters.Get("wei").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "1024000000000000000")
}
