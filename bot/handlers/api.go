package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func SendMessage(id int64, message string) (bool, string) {
	msg := tgbotapi.NewMessage(id, message)
	_, err := HandlerBot.Send(msg)
	if err != nil {
		log.Println(err)
		return false, err.Error()
	}
	return true, ""
}
