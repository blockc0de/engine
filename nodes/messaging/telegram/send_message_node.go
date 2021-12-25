package telegram

import (
	"context"
	"fmt"
	"reflect"

	"github.com/blockc0de/engine/attributes"
	"github.com/blockc0de/engine/block"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	sendMessageNodeDefinition       = []interface{}{attributes.NodeDefinition{NodeName: "SendTelegramMessageNode", FriendlyName: "Send Telegram Message", NodeType: attributes.NodeTypeEnumFunction, GroupName: "Telegram", BlockLimitPerGraph: -1}}
	sendMessageNodeGraphDescription = []interface{}{attributes.NodeGraphDescription{Description: "Send a message on Telegram"}}
)

type SendMessageNode struct {
	block.NodeBase
}

func NewSendMessageNode(id string, graph *block.Graph) (block.Node, error) {
	node := new(SendMessageNode)
	node.NodeData = block.NewNodeData(id, node, graph, reflect.TypeOf(node).String())

	telegramBot, err := block.NewNodeParameter(node, "telegramBot", block.NodeParameterTypeEnumObject, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(telegramBot)

	chatId, err := block.NewNodeParameter(node, "chatId", block.NodeParameterTypeEnumDecimal, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(chatId)

	message, err := block.NewNodeParameter(node, "message", block.NodeParameterTypeEnumString, true, nil)
	if err != nil {
		return nil, err
	}
	node.NodeData.InParameters.Append(message)

	return node, nil
}

func (n *SendMessageNode) CanExecute() bool {
	return true
}

func (n *SendMessageNode) CanBeExecuted() bool {
	return true
}

func (n *SendMessageNode) GetCustomAttributes(t reflect.Type) []interface{} {
	switch t {
	case reflect.TypeOf(attributes.NodeDefinition{}):
		return sendMessageNodeDefinition
	case reflect.TypeOf(attributes.NodeGraphDescription{}):
		return sendMessageNodeGraphDescription
	default:
		return nil
	}
}

func (n *SendMessageNode) OnExecution(ctx context.Context, scheduler block.NodeScheduler) error {
	value := n.Data().InParameters.Get("telegramBot").ComputeValue()
	if value == nil {
		return block.ErrInvalidParameter{Name: "telegramBot"}
	}
	botInstanceNode, ok := value.(*BotInstanceNode)
	if !ok {
		return block.ErrInvalidParameter{Name: "telegramBot"}
	}

	var converter block.NodeParameterConverter
	chatId, ok := converter.ToDecimal(n.Data().InParameters.Get("chatId").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "chatId"}
	}

	message, ok := converter.ToString(n.Data().InParameters.Get("message").ComputeValue())
	if !ok {
		return block.ErrInvalidParameter{Name: "message"}
	}

	_, err := botInstanceNode.bot.Send( tgbotapi.NewMessage(chatId.IntPart(), message))
	if err != nil {
		scheduler.AppendLog("warn", fmt.Sprintf("Failed to send telegram message, reason: %s", err.Error()))
	}

	return nil
}
