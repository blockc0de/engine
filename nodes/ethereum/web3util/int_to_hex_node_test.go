package web3util

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestIntegerToHexNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	n, err := vars.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	d, _ := decimal.NewFromString("212131111111111111111111111111")
	n.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: d}

	integerToHex, err := NewIntegerToHexNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	integerToHex.Data().InParameters.Get("integer").Assignments = n.Data().OutParameters.Get("value")

	result := integerToHex.ComputeParameterValue(integerToHex.Data().OutParameters.Get("hex").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "0x2ad6ebe0febe95fb2231c71c7")
}
