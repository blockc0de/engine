package loader

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blockc0de/engine/block"
	"github.com/shopspring/decimal"
)

func unmarshalNodeParameterValue(valueType block.NodeParameterTypeEnum, data []byte) (interface{}, error) {
	if data == nil || bytes.Compare(data, []byte("null")) == 0 {
		return nil, nil
	}

	switch valueType {
	case block.NodeParameterTypeEnumNode:
		var s string
		err := json.Unmarshal(data, &s)
		return s, err
	case block.NodeParameterTypeEnumBool:
		b := block.NodeParameterBool(false)
		err := json.Unmarshal(data, &b)
		return b, err
	case block.NodeParameterTypeEnumString:
		s := block.NodeParameterString("")
		err := json.Unmarshal(data, &s)
		return s, err
	case block.NodeParameterTypeEnumDecimal:
		d := block.NodeParameterDecimal{Decimal: decimal.Zero}
		err := json.Unmarshal(data, &d)
		return d, err
	case block.NodeParameterTypeEnumStream:
		return nil, nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid %s value", valueType))
	}
}

type NodeParameterSchema struct {
	NodeParameterBaseSchema
	Value interface{} `json:"-"`
}

type NodeParameterBaseSchema struct {
	node             block.Node                  `json:"-"`
	nodeParameter    *block.NodeParameter        `json:"-"`
	Id               string                      `json:"id"`
	Name             string                      `json:"name"`
	Type             block.NodeParameterTypeEnum `json:"type"`
	RawValue         json.RawMessage             `json:"value"`
	Assignment       string                      `json:"assignment"`
	AssignmentNode   string                      `json:"assignment_node"`
	ValueIsReference bool                        `json:"value_is_reference"`
}

func NewNodeParameterSchema(node block.Node, nodeParameter *block.NodeParameter) NodeParameterSchema {
	schema := NodeParameterSchema{
		NodeParameterBaseSchema: NodeParameterBaseSchema{
			node:          node,
			nodeParameter: nodeParameter,
			Id:            nodeParameter.Id,
			Name:          nodeParameter.Name,
			Type:          nodeParameter.ValueType,
		},
	}

	if !nodeParameter.IsDynamic {
		if nodeParameter.ValueType == block.NodeParameterTypeEnumNode {
			if nodeParameter.Value != nil {
				schema.Value = (nodeParameter.Value.(block.Node)).Data().Id
			}
			schema.ValueIsReference = true
		} else {
			schema.ValueIsReference = false
			schema.Value = nodeParameter.Value
		}
	}

	if nodeParameter.Assignments != nil {
		schema.Assignment = nodeParameter.Assignments.Id
		assignmentNode := node.Data().Graph.FindOutParameterNode(nodeParameter.Assignments.Id)
		if assignmentNode != nil {
			schema.AssignmentNode = assignmentNode.Data().Id
		}
	}

	if schema.Value == nil {
		schema.RawValue = json.RawMessage("null")
	} else {
		schema.RawValue, _ = json.Marshal(schema.Value)
	}
	return schema
}

func (s *NodeParameterSchema) UnmarshalJSON(data []byte) (err error) {
	if err = json.Unmarshal(data, &s.NodeParameterBaseSchema); err != nil {
		return err
	}

	s.Value, err = unmarshalNodeParameterValue(s.NodeParameterBaseSchema.Type, s.NodeParameterBaseSchema.RawValue)
	return err
}
