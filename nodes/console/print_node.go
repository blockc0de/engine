package console

import (
	"context"
	"errors"
	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	"math/big"
	"reflect"
)

var (
	printNodeDefinition                = []interface{}{attributes.NodeDefinition{NodeName: "PrintNode", FriendlyName: "Print", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Log", BlockLimitPerGraph: -1}}
	printNodeGraphDescription          = []interface{}{attributes.NodeGraphDescription{Description: "Display a message in the console logs"}}
	printNodeGraphNodeGasConfiguration = []interface{}{attributes.NodeGasConfiguration{BlockGasPrice: big.NewInt(10000000000000)}}
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
	case reflect.TypeOf(attributes.NodeGasConfiguration{}):
		return printNodeGraphNodeGasConfiguration
	default:
		return nil
	}
}

func (n *PrintNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	message := n.Data().InParameters.Get("message")
	messageVal, ok := converter.ToString(message.ComputeValue())
	if !ok {
		return errors.New("invalid message")
	}

	scheduler.AppendLog("info", messageVal)
	return nil
}
