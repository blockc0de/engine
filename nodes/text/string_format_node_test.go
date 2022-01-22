package text

import (
	"testing"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes/vars"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStringFormatNode(t *testing.T) {
	graph := block.NewGraph("", "test")

	format, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	format.Data().OutParameters.Get("value").Value = block.NodeParameterString("Hello {0}, we are greeting you here: {1}! {2}")

	args, err := vars.NewStringNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	args.Data().OutParameters.Get("value").Value = block.NodeParameterString(`["Bob", "blockc0de", 2022]`)

	stringFormat, err := NewStringFormatNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	stringFormat.Data().InParameters.Get("format").Assignments = format.Data().OutParameters.Get("value")
	stringFormat.Data().InParameters.Get("args").Assignments = args.Data().OutParameters.Get("value")

	result := stringFormat.ComputeParameterValue(stringFormat.Data().OutParameters.Get("string").Id, nil)
	assert.Equal(t, string(result.(block.NodeParameterString)), "Hello Bob, we are greeting you here: blockc0de! 2022")
}
