package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

var (
	HandlerUpdate  tgbotapi.Update
	HandlerBot     *tgbotapi.BotAPI
	chatMessage    = tgbotapi.NewMessage(0, "")
	channelMessage = tgbotapi.NewMessage(0, "")
)

func messageHandler() {
	chatMessage.Text = strconv.Itoa(HandlerUpdate.Message.From.ID)
	chatMessage.ReplyToMessageID = HandlerUpdate.Message.MessageID
}

func errorHandler() {
	chatMessage.Text = "I do not know that command"
}

func ChatHandler() {
	switch HandlerUpdate.Message.Command() {
	case "getID":
		messageHandler()
	default:
		errorHandler()
	}
	chatMessage.ChatID = HandlerUpdate.Message.Chat.ID
	_, _ = HandlerBot.Send(chatMessage)
}
