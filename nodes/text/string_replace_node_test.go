package text

import (
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringReplaceNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	original, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	original.Data().OutParameters.Get("value").Value = block.NodeParameterString("aaaaabc")

	toReplace, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	toReplace.Data().OutParameters.Get("value").Value = block.NodeParameterString("aaaaa")

	replaceText, err := variable.NewStringNode(uuid.New().String(), graph)
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
