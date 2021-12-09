package loader

import (
	"errors"
	"github.com/graphlinq-go/engine/block"
	"github.com/graphlinq-go/engine/nodes"
	"github.com/graphlinq-go/engine/nodes/console"
	"github.com/graphlinq-go/engine/nodes/functions"
	"github.com/graphlinq-go/engine/nodes/math"
	"github.com/graphlinq-go/engine/nodes/text"
	"github.com/graphlinq-go/engine/nodes/variable"
	"reflect"
)

type nodeCreator struct {
	Name    string
	Creator func(id string, graph *block.Graph) (block.Node, error)
}

var (
	nodeCreators = []nodeCreator{
		// Common
		{reflect.TypeOf(new(nodes.EntryPointNode)).String(), nodes.NewEntryPointNode},
		{reflect.TypeOf(new(nodes.OnGraphStartNode)).String(), nodes.NewOnGraphStartNode},
		{reflect.TypeOf(new(nodes.StopGraphNode)).String(), nodes.NewStopGraphNode},

		// Base Variable
		{reflect.TypeOf(new(variable.StringNode)).String(), variable.NewStringNode},
		{reflect.TypeOf(new(variable.DecimalNode)).String(), variable.NewDecimalNode},
		{reflect.TypeOf(new(variable.BoolNode)).String(), variable.NewBoolNode},

		// Math
		{reflect.TypeOf(new(math.AddNode)).String(), math.NewAddNode},
		{reflect.TypeOf(new(math.SubNode)).String(), math.NewSubNode},
		{reflect.TypeOf(new(math.MulNode)).String(), math.NewMulNode},
		{reflect.TypeOf(new(math.DivNode)).String(), math.NewDivNode},
		{reflect.TypeOf(new(math.PowNode)).String(), math.NewPowNode},
		{reflect.TypeOf(new(math.ModNode)).String(), math.NewModNode},
		{reflect.TypeOf(new(math.FloorNode)).String(), math.NewFloorNode},
		{reflect.TypeOf(new(math.CeilNode)).String(), math.NewCeilNode},
		{reflect.TypeOf(new(math.RoundNode)).String(), math.NewRoundNode},
		{reflect.TypeOf(new(math.TruncNode)).String(), math.NewTruncNode},
		{reflect.TypeOf(new(math.PercentageDiffNode)).String(), math.NewPercentageDiffNode},

		// String
		{reflect.TypeOf(new(text.StringConcatNode)).String(), text.NewStringConcatNode},
		{reflect.TypeOf(new(text.StringReplaceNode)).String(), text.NewStringReplaceNode},
		{reflect.TypeOf(new(text.StringContainsNode)).String(), text.NewStringContainsNode},
		{reflect.TypeOf(new(text.StringToLowerNode)).String(), text.NewStringToLowerNode},
		{reflect.TypeOf(new(text.StringToUpperNode)).String(), text.NewStringToUpperNode},
		{reflect.TypeOf(new(text.StringTrimLeftNode)).String(), text.NewStringTrimLeftNode},
		{reflect.TypeOf(new(text.StringTrimRightNode)).String(), text.NewStringTrimRightNode},

		// Log
		{reflect.TypeOf(new(console.PrintNode)).String(), console.NewPrintNode},

		// Function
		{reflect.TypeOf(new(functions.FunctionNode)).String(), functions.NewFunctionNode},
	}

	nodeCreatorsMapper = make(map[string]func(id string, graph *block.Graph) (block.Node, error))
)

func init() {
	for _, item := range nodeCreators {
		nodeCreatorsMapper[item.Name] = item.Creator
	}
}

type NodeFactory struct{}

func (NodeFactory) NewNode(nodeType string, id string, graph *block.Graph) (block.Node, error) {
	creator, ok := nodeCreatorsMapper[nodeType]
	if !ok {
		return nil, errors.New("node creator unregistered")
	}
	return creator(id, graph)
}
