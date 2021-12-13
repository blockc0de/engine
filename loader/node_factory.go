package loader

import (
	"errors"
	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes"
	"github.com/blockc0de/engine/nodes/console"
	"github.com/blockc0de/engine/nodes/encoding"
	"github.com/blockc0de/engine/nodes/functions"
	"github.com/blockc0de/engine/nodes/math"
	"github.com/blockc0de/engine/nodes/text"
	"github.com/blockc0de/engine/nodes/time"
	"github.com/blockc0de/engine/nodes/vars"
	"reflect"
)

type nodeCreator struct {
	Name    string
	Creator func(id string, graph *block.Graph) (block.Node, error)
}

// The registry for the node type and creator
var (
	nodeCreators = []nodeCreator{
		// Common
		{reflect.TypeOf(new(nodes.EntryPointNode)).String(), nodes.NewEntryPointNode},
		{reflect.TypeOf(new(nodes.OnGraphStartNode)).String(), nodes.NewOnGraphStartNode},
		{reflect.TypeOf(new(nodes.StopGraphNode)).String(), nodes.NewStopGraphNode},

		// Variable
		{reflect.TypeOf(new(vars.StringNode)).String(), vars.NewStringNode},
		{reflect.TypeOf(new(vars.DecimalNode)).String(), vars.NewDecimalNode},
		{reflect.TypeOf(new(vars.BoolNode)).String(), vars.NewBoolNode},
		{reflect.TypeOf(new(vars.SetVariableNode)).String(), vars.NewSetVariableNode},
		{reflect.TypeOf(new(vars.GetVariableNode)).String(), vars.NewGetVariableNode},
		{reflect.TypeOf(new(vars.IsVariableExistNode)).String(), vars.NewIsVariableExistNode},

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

		// Time
		{reflect.TypeOf(new(time.TimerNode)).String(), time.NewTimerNode},
		{reflect.TypeOf(new(time.GetTimestampNode)).String(), time.NewGetTimestampNode},
		{reflect.TypeOf(new(time.ParseTimestampNode)).String(), time.NewParseTimestampNode},
		{reflect.TypeOf(new(time.FormatTimestampNode)).String(), time.NewFormatTimestampNode},

		// JSON
		{reflect.TypeOf(new(encoding.CreateJsonObjectNode)).String(), encoding.NewCreateJsonObjectNode},
		{reflect.TypeOf(new(encoding.AddJsonValueNode)).String(), encoding.NewAddJsonValueNode},
		{reflect.TypeOf(new(encoding.ConvertToJsonNode)).String(), encoding.NewConvertToJsonNode},
		{reflect.TypeOf(new(encoding.LastNodeToJsonNode)).String(), encoding.NewLastNodeToJsonNode},
		{reflect.TypeOf(new(encoding.JsonToJsonObjectNode)).String(), encoding.NewJsonToJsonObjectNode},

		// Function
		{reflect.TypeOf(new(functions.FunctionNode)).String(), functions.NewFunctionNode},
	}

	nodeCreatorMapper = make(map[string]func(id string, graph *block.Graph) (block.Node, error))
)

func init() {
	for _, item := range nodeCreators {
		nodeCreatorMapper[item.Name] = item.Creator
	}
}

// NewNode create node instance by node type
func NewNode(nodeType string, id string, graph *block.Graph) (block.Node, error) {
	creator, ok := nodeCreatorMapper[nodeType]
	if !ok {
		return nil, errors.New("node unregistered")
	}
	return creator(id, graph)
}

// RegisterNodeType register a new node type, use reflection to create a node when loading a graph
func RegisterNodeType(nodeType string, creator func(id string, graph *block.Graph) (block.Node, error)) error {
	creator, ok := nodeCreatorMapper[nodeType]
	if !ok {
		return errors.New("node already registered")
	}

	nodeCreatorMapper[nodeType] = creator
	nodeCreators = append(nodeCreators, nodeCreator{
		Name:    nodeType,
		Creator: creator,
	})
	return nil
}
