package block

import (
	"errors"
	"github.com/google/uuid"
)

type NodeParameter struct {
	Id          string
	Node        Node `json:"-"'`
	Name        string
	Value       interface{}
	ValueType   NodeParameterTypeEnum
	IsIn        bool
	IsDynamic   bool
	Assignments *NodeParameter
	IsReference bool
}

func NewNodeParameter(
	node Node, name string, valueType NodeParameterTypeEnum, isIn bool, value interface{}) (*NodeParameter, error) {

	if value != nil {
		t := GetNodeParameterTypeEnum(value)
		if t != valueType || t == NodeParameterTypeEnumUnknown {
			return nil, errors.New("invalid type for the value")
		}
	}

	p := NodeParameter{
		Id:          uuid.New().String(),
		Node:        node,
		Name:        name,
		ValueType:   valueType,
		IsIn:        isIn,
		IsReference: valueType == NodeParameterTypeEnumNode,
		Value:       value,
	}
	return &p, nil
}

func NewDynamicNodeParameter(node Node, name string, valueType NodeParameterTypeEnum, isIn bool) (*NodeParameter, error) {
	p := NodeParameter{
		Id:          "",
		Node:        node,
		Name:        name,
		ValueType:   valueType,
		IsIn:        isIn,
		IsDynamic:   true,
		IsReference: valueType == NodeParameterTypeEnumNode,
	}

	return &p, nil
}

func (p *NodeParameter) ComputeValue() (value interface{}) {
	defer func() {
		if err := recover(); err != nil {
			value = nil
		}
	}()

	if p.IsIn && p.Assignments != nil {
		value = p.Node.ComputeParameterValue(p.Id, p.Assignments.ComputeValue())
		if value == nil && p.ValueType == NodeParameterTypeEnumString {
			value = ""
		}
		return value
	}
	return p.Node.ComputeParameterValue(p.Id, p.Value)
}

type NodeParameters []*NodeParameter

func (p *NodeParameters) Get(name string) *NodeParameter {
	for _, parameter := range *p {
		if parameter.Name == name {
			return parameter
		}
	}
	return nil
}

func (p *NodeParameters) Append(parameter *NodeParameter) {
	*p = append(*p, parameter)
}
