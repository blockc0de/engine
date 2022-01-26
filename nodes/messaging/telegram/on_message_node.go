package telegram

import (
	"context"
	"reflect"

	"github.com/shopspring/decimal"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	onMessageTelegramBotNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "OnMessageTelegramBotNode", FriendlyName: "On Telegram Message", NodeType: attributes.NodeTypeEnumEvent, GroupName: "Telegram"}}
	onMessageTelegramBotNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "This event allow you to handle a message from Telegram in your grap"}}
)

type OnMessageTelegramBotNode struct {
	block.NodeBase
}

func NewOnMessageTelegramBotNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(OnMessageTelegramBotNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())
	node.NodeData.IsEventNode = true

	telegramBot, err := block.NewNodeParameter(node, "telegramBot", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(telegramBot)

	from, err := block.NewDynamicNodeParameter(node, "from", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(from)

	fromId, err := block.NewDynamicNodeParameter(node, "fromId", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(fromId)

	chatId, err := block.NewDynamicNodeParameter(node, "chatId", block.NodeParameterTypeEnumDecimal, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(chatId)

	message, err := block.NewDynamicNodeParameter(node, "message", block.NodeParameterTypeEnumString, false)
	if err != nil {
		return nil, err
	}
	node.NodeData.OutParameters.Append(message)

	return node, nil
}

func (n *OnMessageTelegramBotNode) CanExecute() bool {
	return true
}

func (n *OnMessageTelegramBotNode) pollingUpdates(bot *tgbotapi.BotAPI, engine block.Engine) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ch := bot.GetUpdatesChan(u)

	for update := range ch {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat == nil || update.Message.From == nil {
			continue
		}

		parameters := make(block.NodeParameters, 0, 4)
		from, err := block.NewDynamicNodeParameter(n, "from", block.NodeParameterTypeEnumString, false)
		if err != nil {
			continue
		}
		from.Value = block.NodeParameterString(update.Message.From.UserName)
		parameters.Append(from)

		fromId, err := block.NewDynamicNodeParameter(n, "fromId", block.NodeParameterTypeEnumDecimal, false)
		if err != nil {
			continue

		}
		fromId.Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(update.Message.From.ID)}
		parameters.Append(fromId)

		chatId, err := block.NewDynamicNodeParameter(n, "chatId", block.NodeParameterTypeEnumDecimal, false)
		if err != nil {
			continue
		}
		chatId.Value = block.NodeParameterDecimal{Decimal: decimal.NewFromInt(update.Message.Chat.ID)}
		parameters.Append(chatId)

		message, err := block.NewDynamicNodeParameter(n, "message", block.NodeParameterTypeEnumString, false)
		if err != nil {
			continue
		}
		message.Value = block.NodeParameterString(update.Message.Text)
		parameters.Append(message)

		engine.AddCycle(n, parameters)
	}
}

func (n *OnMessageTelegramBotNode) SetupEvent(engine block.Engine) error {
	value := n.Data().InParameters.Get("telegramBot").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "telegramBot"}
	}
	botInstanceNode, ok := value.(*BotInstanceNode)
	if !ok {
		return block.ErrInvalidParameter{Name: "telegramBot"}
	}

	go n.pollingUpdates(botInstanceNode.bot, engine)

	return nil
}

func (n *OnMessageTelegramBotNode) BeginCycle(ctx context.Context, engine block.Engine) {
	engine.NextNode(ctx, n)
}

func (n *OnMessageTelegramBotNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return onMessageTelegramBotNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return onMessageTelegramBotNodeGraphDescription
	default:
		return nil
	}
}

func (n *OnMessageTelegramBotNode) OnExecution(context.Context, block.Engine) error {
	return nil
}

func (n *OnMessageTelegramBotNode) OnStop() error {
	return nil
}
