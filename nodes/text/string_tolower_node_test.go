package text

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToLowerNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	s, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	s.Data().OutParameters.Get("value").Value = block.NodeParameterString("ABCDE")

	stringToLower, err := NewStringToLowerNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringToLower.Data().InParameters.Get("input").Assignments = s.Data().OutParameters.Get("value")

	result := stringToLower.ComputeParameterValue(stringToLower.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "abcde")
}
