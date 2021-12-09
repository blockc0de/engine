package interop

import (
	"github.com/graphlinq-go/engine/block"
	jsoniter "github.com/json-iterator/go"
)

type GraphSchema struct {
	graph *block.Graph `json:"-"`
	Name  string       `json:"name"`
	Nodes []NodeSchema `json:"nodes"`
}

func NewGraphSchema(graph *block.Graph) GraphSchema {
	return GraphSchema{
		graph: graph,
		Nodes: make([]NodeSchema, 0),
	}
}

func (schema *GraphSchema) Export() ([]byte, error) {
	schema.Name = schema.graph.Name
	for _, node := range schema.graph.Nodes {
		nodeSchema := NewNodeSchema(node)
		schema.Nodes = append(schema.Nodes, nodeSchema)
	}
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(schema)
}
