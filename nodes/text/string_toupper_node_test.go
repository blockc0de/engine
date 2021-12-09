package text

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes/variable"
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
