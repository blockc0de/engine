package interop

import (
	"errors"
	"reflect"

	"github.com/blockc0de/engine/block"
	"github.com/blockc0de/engine/nodes"
	"github.com/blockc0de/engine/nodes/array"
	"github.com/blockc0de/engine/nodes/branch"
	"github.com/blockc0de/engine/nodes/console"
	"github.com/blockc0de/engine/nodes/encoding"
	"github.com/blockc0de/engine/nodes/ethereum"
	"github.com/blockc0de/engine/nodes/ethereum/web3util"
	"github.com/blockc0de/engine/nodes/functions"
	"github.com/blockc0de/engine/nodes/math"
	"github.com/blockc0de/engine/nodes/messaging/telegram"
	"github.com/blockc0de/engine/nodes/storage"
	"github.com/blockc0de/engine/nodes/text"
	"github.com/blockc0de/engine/nodes/time"
	"github.com/blockc0de/engine/nodes/vars"
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

		// Condition
		{reflect.TypeOf(new(branch.BoolBranchNode)).String(), branch.NewBoolBranchNode},
		{reflect.TypeOf(new(branch.StringBranchNode)).String(), branch.NewStringBranchNode},
		{reflect.TypeOf(new(branch.DecimalBranchNode)).String(), branch.NewDecimalBranchNode},
		{reflect.TypeOf(new(branch.DecimalRangeBranchNode)).String(), branch.NewDecimalRangeBranchNode},

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

		// Time
		{reflect.TypeOf(new(time.TimerNode)).String(), time.NewTimerNode},
		{reflect.TypeOf(new(time.GetTimestampNode)).String(), time.NewGetTimestampNode},
		{reflect.TypeOf(new(time.ParseTimestampNode)).String(), time.NewParseTimestampNode},
		{reflect.TypeOf(new(time.FormatTimestampNode)).String(), time.NewFormatTimestampNode},

		// JSON
		{reflect.TypeOf(new(encoding.MergeJsonNodeNode)).String(), encoding.NewMergeJsonNodeNode},
		{reflect.TypeOf(new(encoding.JsonSelectorNodeNode)).String(), encoding.NewJsonSelectorNodeNode},
		{reflect.TypeOf(new(encoding.CreateJsonObjectNode)).String(), encoding.NewCreateJsonObjectNode},
		{reflect.TypeOf(new(encoding.AddJsonValueNode)).String(), encoding.NewAddJsonValueNode},
		{reflect.TypeOf(new(encoding.ConvertToJsonNode)).String(), encoding.NewConvertToJsonNode},
		{reflect.TypeOf(new(encoding.LastNodeToJsonNode)).String(), encoding.NewLastNodeToJsonNode},
		{reflect.TypeOf(new(encoding.JsonToJsonObjectNode)).String(), encoding.NewJsonToJsonObjectNode},
		{reflect.TypeOf(new(encoding.JsonDeserializeToArrayNode)).String(), encoding.NewJsonDeserializeToArrayNode},

		// Array
		{reflect.TypeOf(new(array.CreateArrayNode)).String(), array.NewCreateArrayNode},
		{reflect.TypeOf(new(array.GetArraySizeNode)).String(), array.NewGetArraySizeNode},
		{reflect.TypeOf(new(array.AddElementNode)).String(), array.NewAddElementNode},
		{reflect.TypeOf(new(array.GetElementAtIndexNode)).String(), array.NewGetElementAtIndexNode},
		{reflect.TypeOf(new(array.EachElementArrayNode)).String(), array.NewEachElementArrayNode},
		{reflect.TypeOf(new(array.ClearArrayNode)).String(), array.NewClearArrayNode},

		// Log
		{reflect.TypeOf(new(console.PrintNode)).String(), console.NewPrintNode},

		// Function
		{reflect.TypeOf(new(functions.FunctionNode)).String(), functions.NewFunctionNode},
		{reflect.TypeOf(new(functions.CreateFunctionParametersNode)).String(), functions.NewCreateFunctionParametersNode},
		{reflect.TypeOf(new(functions.AdaddFunctionParameterNode)).String(), functions.NewAdaddFunctionParameterNode},
		{reflect.TypeOf(new(functions.CallFunctionNode)).String(), functions.NewCallFunctionNode},

		// Storage
		{reflect.TypeOf(new(storage.SetWalletKeyNode)).String(), storage.NewSetWalletKeyNode},
		{reflect.TypeOf(new(storage.GetWalletKeyNode)).String(), storage.NewGetWalletKeyNode},
		{reflect.TypeOf(new(storage.ListWalletKeyNode)).String(), storage.NewListWalletKeyNode},
		{reflect.TypeOf(new(storage.DelWalletKeyNode)).String(), storage.NewDelWalletKeyNode},

		// Web3.Util
		{reflect.TypeOf(new(web3util.HexToIntegerNode)).String(), web3util.NewHexToIntegerNode},
		{reflect.TypeOf(new(web3util.IntegerToHexNode)).String(), web3util.NewIntegerToHexNode},
		{reflect.TypeOf(new(web3util.FromWeiNode)).String(), web3util.NewFromWeiNode},
		{reflect.TypeOf(new(web3util.ToWeiNode)).String(), web3util.NewToWeiNode},

		// Blockchain.Ethereum
		{reflect.TypeOf(new(ethereum.EthConnection)).String(), ethereum.NewEthConnection},
		{reflect.TypeOf(new(ethereum.GetBalanceNode)).String(), ethereum.NewGetBalanceNode},
		{reflect.TypeOf(new(ethereum.GetErc20BalanceNode)).String(), ethereum.NewGetErc20BalanceNode},
		{reflect.TypeOf(new(ethereum.GetErc20TokenInfoNode)).String(), ethereum.NewGetErc20TokenInfoNode},
		{reflect.TypeOf(new(ethereum.OnNewBlockEventNode)).String(), ethereum.NewOnNewBlockEventNode},
		{reflect.TypeOf(new(ethereum.OnEventLogNode)).String(), ethereum.NewOnEventLogNode},
		{reflect.TypeOf(new(ethereum.CallContractNode)).String(), ethereum.NewCallContractNode},
		{reflect.TypeOf(new(ethereum.SendTransactionNode)).String(), ethereum.NewSendTransactionNode},
		{reflect.TypeOf(new(ethereum.JsonAbiNode)).String(), ethereum.NewJsonAbiNode},
		{reflect.TypeOf(new(ethereum.Erc20AbiNode)).String(), ethereum.NewErc20AbiNode},
		{reflect.TypeOf(new(ethereum.CallResultDecoderNode)).String(), ethereum.NewCallResultDecoderNode},
		{reflect.TypeOf(new(ethereum.AbiEncoderNode)).String(), ethereum.NewAbiEncoderNode},
		{reflect.TypeOf(new(ethereum.AbiDecoderNode)).String(), ethereum.NewAbiDecoderNode},
		{reflect.TypeOf(new(ethereum.LogDecoderNode)).String(), ethereum.NewLogDecoderNode},

		// Telegram
		{reflect.TypeOf(new(telegram.BotInstanceNode)).String(), telegram.NewBotInstanceNode},
		{reflect.TypeOf(new(telegram.OnMessageTelegramBotNode)).String(), telegram.NewOnMessageTelegramBotNode},
		{reflect.TypeOf(new(telegram.SendMessageNode)).String(), telegram.NewSendMessageNode},
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
	_, ok := nodeCreatorMapper[nodeType]
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
