package loader

import (
	"github.com/blockc0de/engine/block"
)

type NodeSchema struct {
	node          block.Node            `json:"-"`
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	OutNode       *string               `json:"out_node"`
	CanBeExecuted bool                  `json:"can_be_executed"`
	CanExecute    bool                  `json:"can_execute"`
	InParameters  []NodeParameterSchema `json:"in_parameters"`
	OutParameters []NodeParameterSchema `json:"out_parameters"`
}

func NewNodeSchema(node block.Node) NodeSchema {
	schema := NodeSchema{
		node:          node,
		Id:            node.Data().Id,
		Type:          node.Data().NodeType,
		CanBeExecuted: node.CanBeExecuted(),
		CanExecute:    node.CanExecute(),
	}

	if node.Data().OutNode != nil {
		id := node.Data().OutNode.Data().Id
		schema.OutNode = &id
	}

	schema.InParameters = make([]NodeParameterSchema, 0)
	for _, parameter := range node.Data().InParameters {
		schema.InParameters = append(schema.InParameters, NewNodeParameterSchema(node, parameter))
	}

	schema.OutParameters = make([]NodeParameterSchema, 0)
	for _, parameter := range node.Data().OutParameters {
		schema.OutParameters = append(schema.OutParameters, NewNodeParameterSchema(node, parameter))
	}

	return schema
}
