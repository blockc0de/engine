package text

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/variable"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringTrimRightNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	s, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	s.Data().OutParameters.Get("value").Value = block.NodeParameterString("abc00000")

	cutset, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	cutset.Data().OutParameters.Get("value").Value = block.NodeParameterString("0")

	stringTrimRight, err := NewStringTrimRightNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringTrimRight.Data().InParameters.Get("input").Assignments = s.Data().OutParameters.Get("value")
	stringTrimRight.Data().InParameters.Get("cutset").Assignments = cutset.Data().OutParameters.Get("value")

	result := stringTrimRight.ComputeParameterValue(stringTrimRight.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "abc")
}
