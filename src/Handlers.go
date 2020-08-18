package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func messageHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!")
	//if _, err := api.Send(msg); err != nil {
	//	log.Panic(err)
	//}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, strconv.Itoa(update.Message.From.ID))
	msg.ReplyToMessageID = update.Message.MessageID
	_, _ = bot.Send(msg)
}
