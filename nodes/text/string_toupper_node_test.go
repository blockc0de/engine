package text

import (
	"fmt"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/variable"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToUpperNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	s, err := variable.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	s.Data().OutParameters.Get("value").Value = block.NodeParameterString("abcde")

	stringToUpper, err := NewStringToUpperNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringToUpper.Data().InParameters.Get("input").Assignments = s.Data().OutParameters.Get("value")

	result := stringToUpper.ComputeParameterValue(stringToUpper.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "ABCDE")
	fmt.Println(stringToUpper.Data().OutParameters.Get("string").ComputeValue())
}
