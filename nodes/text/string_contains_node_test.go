package text

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStringContainsNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	s, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	s.Data().OutParameters.Get("value").Value = block.NodeParameterString("abcdefg")

	toSearch, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	toSearch.Data().OutParameters.Get("value").Value = block.NodeParameterString("abc")

	stringContains, err := NewStringContainsNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	assert.True(t, stringContains.CanBeExecuted())
	stringContains.Data().InParameters.Get("string").Assignments = s.Data().OutParameters.Get("value")
	stringContains.Data().InParameters.Get("toSearch").Assignments = toSearch.Data().OutParameters.Get("value")
}
