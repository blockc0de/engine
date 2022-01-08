package telegram

import (
	"context"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	botInstanceNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "TelegramBotInstanceNode", FriendlyName: "Telegram Bot", NodeType: attributes.NodeTypeEnumConnector, GroupName: "Telegram"}}
	botInstanceNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Telegram connector, to retrieve your AccessToken talk to BotFather to create a new bot"}}
)

type BotInstanceNode struct {
	block.NodeBase
	bot *tgbotapi.BotAPI
}

func NewBotInstanceNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(BotInstanceNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.CanBeSerialized = false

	accessToken, err := block.NewNodeParameter(node, "accessToken", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(accessToken)

	telegramBot, err := block.NewDynamicNodeParameter(node, "telegramBot", block.NodeParameterTypeEnumObject, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(telegramBot)

	username, err := block.NewDynamicNodeParameter(node, "username", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(username)

	return node, nil
}

func (n *BotInstanceNode) CanExecute() bool {
	return true
}

func (n *BotInstanceNode) SetupConnector(scheduler block.NodeScheduler) error {
	var converter block.NodeParameterConverter
	accessToken, ok := converter.ToString(n.Data().InParameters.Get("accessToken").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "accessToken"}
	}

	var err error
	n.bot, err = tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		return err
	}

	n.Data().OutParameters.Get("username").Value = block.NodeParameterString(n.bot.Self.UserName)

	scheduler.AddCycle(n, nil)
	return nil
}

func (n *BotInstanceNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return botInstanceNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return botInstanceNodeGraphDescription
	default:
		return nil
	}
}

func (n *BotInstanceNode) BeginCycle(ctx context.Context, scheduler block.NodeScheduler) {
	scheduler.NextNode(ctx, n)
}

func (n *BotInstanceNode) ComputeParameterValue(parameterId string, value interface{}) interface{} {
	if parameterId == n.Data().OutParameters.Get("telegramBot").Id {
		return n
	}
	return value
}

func (n *BotInstanceNode) OnExecution(context.Context, block.NodeScheduler) error {
	return nil
}

func (n *BotInstanceNode) OnStop() error {
	if n.bot != nil {
		n.bot.StopReceivingUpdates()
	}

	return nil
}
