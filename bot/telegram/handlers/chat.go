package handlers

import (
	"appbrickie/api/database"
	"appbrickie/bot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	HandlerUpdate  tgbotapi.Update
	HandlerBot     *tgbotapi.BotAPI
	chatMessage    = tgbotapi.NewMessage(0, "")
	channelMessage = tgbotapi.NewMessage(0, "")
)

func startHandler() {
	chatMessage.Text = utils.StartTemplate(HandlerUpdate.Message.Chat.FirstName)
}

func idGenerator() {
	user, resp := database.ServiceHelper.FetchUser(HandlerUpdate.Message.Chat.ID)
	if !resp {
		uid, success := database.ServiceHelper.CreateUser(HandlerUpdate.Message.Chat.ID, HandlerUpdate.Message.Chat.UserName, HandlerUpdate.Message.Chat.FirstName+" "+HandlerUpdate.Message.Chat.LastName, false)
		if success {
			chatMessage.Text = uid
		} else {
			chatMessage.Text = "Error Generating User ID , please try again later."
		}
		chatMessage.ReplyToMessageID = HandlerUpdate.Message.MessageID
	} else {
		chatMessage.Text = user.UniqueId
		chatMessage.ReplyToMessageID = HandlerUpdate.Message.MessageID
	}
}

func errorHandler() {
	chatMessage.Text = "I do not know that command"
}

func greetHandler() {
	chatMessage.Text = "Hello " + HandlerUpdate.Message.Chat.FirstName + "!"
}
func helpHandler() {
	chatMessage.Text = utils.GetHelpTemplate(HandlerUpdate.Message.Chat.FirstName)
}

func ChatHandler() {
	switch HandlerUpdate.Message.Command() {
	case "start":
		startHandler()
	case "getid":
		idGenerator()
	case "greet":
		greetHandler()
	case "help":
		helpHandler()
	default:
		errorHandler()
	}
	chatMessage.ChatID = HandlerUpdate.Message.Chat.ID
	_, _ = HandlerBot.Send(chatMessage)
	chatMessage.ReplyToMessageID = -1
}
