package web3util

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHexToIntegerNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	hex, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	hex.Data().OutParameters.Get("value").Value = block.NodeParameterString("0x0000000000000000000000000000000000000002ad6ebe0febe95fb2231c71c7")

	hexToInteger, err := NewHexToIntegerNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	hexToInteger.Data().InParameters.Get("hex").Assignments = hex.Data().OutParameters.Get("value")

	result := hexToInteger.ComputeParameterValue(hexToInteger.Data().OutParameters.Get("integer").Id, nil)
	assert.Equal(t, result.(block.NodeParameterDecimal).String(), "212131111111111111111111111111")
}
