package block

import (
	"context"
	"reflect"
	"sort"
)

type Node interface {
	Data() *NodeData
	CanExecute() bool
	CanBeExecuted() bool
	GetCustomAttributes(reflect.Type) []interface{}
	ComputeParameterValue(string, interface{}) interface{}
	OnStop() error
}

type RootNode interface {
	Node
	BeginCycle(context.Context, NodeExecutor) error
}

type EventNode interface {
	Node
	SetupEvent(context.Context, NodeExecutor) error
}

type ConnectorNode interface {
	Node
	SetupConnector(context.Context, NodeExecutor) error
}

type ExecutableNode interface {
	Node
	OnExecution(context.Context, NodeExecutor) error
}

// NodeBase Base class for all nodes
type NodeBase struct {
	NodeData *NodeData
}

func (n *NodeBase) Data() *NodeData {
	return n.NodeData
}

func (n *NodeBase) CanExecute() bool {
	return false
}

func (n *NodeBase) CanBeExecuted() bool {
	return false
}

func (n *NodeBase) GetCustomAttributes(reflect.Type) []interface{} {
	return nil
}

func (n *NodeBase) ComputeParameterValue(id string, value interface{}) interface{} {
	return value
}

func (n *NodeBase) OnStop() error {
	return nil
}

// Nodes The node slice of sort interface is implemented
// event node is at the front of the slice.
type Nodes []Node

func (x Nodes) Len() int { return len(x) }

func (x Nodes) Less(i, j int) bool {
	if x[i].Data().IsEventNode && x[j].Data().IsEventNode {
		return i < j
	}
	if x[i].Data().IsEventNode {
		return true
	}
	if x[j].Data().IsEventNode {
		return false
	}
	return i < j
}

func (x Nodes) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x Nodes) Sort() { sort.Sort(x) }
