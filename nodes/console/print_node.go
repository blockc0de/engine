package console

import (
	"context"
	"fmt"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
)

var (
	printNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "PrintNode", FriendlyName: "Print", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Log"}}
	printNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Display a message in the console logs"}}
)

type PrintNode struct {
	block.NodeBase
}

func NewPrintNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(PrintNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	message, err := block.NewNodeParameter(node, "message", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(message)

	return node, err
}

func (n *PrintNode) CanExecute() bool {
	return true
}

func (n *PrintNode) CanBeExecuted() bool {
	return true
}

func (n *PrintNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return printNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return printNodeGraphDescription
	default:
		return nil
	}
}

func (n *PrintNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	value := n.Data().InParameters.Get("message").ComputeValue()
	message, ok := converter.ToString(value)
	if !ok {
		message = fmt.Sprintf("%+v", value)
	}

	scheduler.AppendLog("info", message)
	return nil
}
