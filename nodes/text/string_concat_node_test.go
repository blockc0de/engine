package text

import (
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringConcatNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterString("abc")

	b, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterString("def")

	delimiter, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	delimiter.Data().OutParameters.Get("value").Value = block.NodeParameterString(",")

	stringConcat, err := NewStringConcatNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringConcat.Data().InParameters.Get("stringA").Assignments = a.Data().OutParameters.Get("value")
	stringConcat.Data().InParameters.Get("stringB").Assignments = b.Data().OutParameters.Get("value")
	stringConcat.Data().InParameters.Get("delimiter").Assignments = delimiter.Data().OutParameters.Get("value")

	result := stringConcat.ComputeParameterValue(stringConcat.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "abc,def")
}
