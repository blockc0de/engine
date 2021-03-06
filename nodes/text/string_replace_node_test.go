package text

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStringReplaceNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	original, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	original.Data().OutParameters.Get("value").Value = block.NodeParameterString("aaaaabc")

	toReplace, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	toReplace.Data().OutParameters.Get("value").Value = block.NodeParameterString("aaaaa")

	replaceText, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	replaceText.Data().OutParameters.Get("value").Value = block.NodeParameterString("11111")

	stringReplace, err := NewStringReplaceNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringReplace.Data().InParameters.Get("original").Assignments = original.Data().OutParameters.Get("value")
	stringReplace.Data().InParameters.Get("toReplace").Assignments = toReplace.Data().OutParameters.Get("value")
	stringReplace.Data().InParameters.Get("replaceText").Assignments = replaceText.Data().OutParameters.Get("value")

	result := stringReplace.ComputeParameterValue(stringReplace.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "11111bc")
}
