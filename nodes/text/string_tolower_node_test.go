package text

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToLowerNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	s, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	s.Data().OutParameters.Get("value").Value = block.NodeParameterString("ABCDE")

	stringToLower, err := NewStringToLowerNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringToLower.Data().InParameters.Get("input").Assignments = s.Data().OutParameters.Get("value")

	result := stringToLower.ComputeParameterValue(stringToLower.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "abcde")
	fmt.Println(stringToLower.Data().OutParameters.Get("string").ComputeValue())
}
