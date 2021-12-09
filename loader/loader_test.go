package loader_test

import (
	"github.com/google/uuid"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/loader"
	"github.com/graphlinq-go/engine/nodes/math"
	"github.com/graphlinq-go/engine/nodes/variable"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	graphJSON []byte
)

func TestExportGraph(t *testing.T) {
	graph := block.NewGraph("", "test")

	a, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	a.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(10)}
	graph.AddNode(a)

	b, err := variable.NewDecimalNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	b.Data().OutParameters.Get("value").Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(20)}
	graph.AddNode(b)

	add, err := math.NewAddNode(uuid.New().String(), graph)
	assert.Nil(t, err)
	graph.AddNode(add)
	add.Data().InParameters.Get("a").Assignments = a.Data().OutParameters.Get("value")
	add.Data().InParameters.Get("b").Assignments = b.Data().OutParameters.Get("value")

	data, err := loader.ExportGraph(graph)
	assert.Nil(t, err)
	graphJSON = data
}

func TestLoadGraph(t *testing.T) {
	graph, err := loader.LoadGraph(graphJSON)
	assert.Nil(t, err)

	data, err := loader.ExportGraph(graph)
	assert.Nil(t, err)

	assert.Equal(t, len(data), len(graphJSON))
}

func TestExportNodeSchema(t *testing.T) {
	_, err := loader.ExportNodeSchema()
	assert.Nil(t, err)
}
