package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
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

func SendPackage(id int64, message string, packagePath string) (bool, error) {
	log.Println(fmt.Println(id, message, packagePath))
	_, err := HandlerBot.UploadFile("sendDocument", map[string]string{
		"chat_id": strconv.Itoa(int(id)),
		"caption": message,
	},
		"document", packagePath)
	if err != nil {
		log.Println("Api SendPackage: " + err.Error())
		SendErrorMessage(id)
		return false, err
	}
	return true, nil
}

func SendErrorMessage(id int64) {
	msg := tgbotapi.NewMessage(id, "There was an error sending your package 😨 Please check your repo for latest build!")
	_, _ = HandlerBot.Send(msg)
}
