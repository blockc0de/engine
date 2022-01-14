package block

import (
	"context"
	"reflect"

	"github.com/go-redis/redis"
)

type Node interface {
	Data() *NodeData
	CanExecute() bool
	CanBeExecuted() bool
	GetCustomAttributes(reflect.Type) []interface{}
	ComputeParameterValue(string, interface{}) interface{}
}

type StartNode interface {
	ExecutableNode
	BeginCycle(context.Context, NodeScheduler)
}

type EventNode interface {
	Node
	SetupEvent(NodeScheduler) error
	OnStop() error
}

type StorageNode interface {
	Node
	SetupDatabase(string, redis.Cmdable) error
}

type ConnectorNode interface {
	Node
	SetupConnector(NodeScheduler) error
	OnStop() error
}

type ExecutableNode interface {
	Node
	OnExecution(context.Context, NodeScheduler) error
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
