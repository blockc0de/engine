package web3util

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestWeiToDecimal(t *testing.T) {
	graph := block.NewGraph("", "test")

	wei, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	wei.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(1024000000000000000)}

	decimals, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	decimals.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(18)}

	fromWeiNode, err := NewFromWeiNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	fromWeiNode.Data().InParameters.Get("wei").Assignments = wei.Data().OutParameters.Get("value")
	fromWeiNode.Data().InParameters.Get("decimals").Assignments = decimals.Data().OutParameters.Get("value")

	result := fromWeiNode.ComputeParameterValue(fromWeiNode.Data().OutParameters.Get("value").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "1.024")
}
