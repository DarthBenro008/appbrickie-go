package handlers

import (
	"appbrickie/api/database"
	"appbrickie/bot/utils"
)

func channelIdGenerator() {
	user, resp := database.ServiceHelper.FetchUser(HandlerUpdate.ChannelPost.Chat.ID)
	if !resp {
		uid, success := database.ServiceHelper.CreateUser(HandlerUpdate.ChannelPost.Chat.ID, HandlerUpdate.ChannelPost.Chat.Title, HandlerUpdate.ChannelPost.Chat.FirstName+" "+HandlerUpdate.ChannelPost.Chat.FirstName, true)
		if success {
			channelMessage.Text = uid
		} else {
			channelMessage.Text = "Error Generating User ID , please try again later."
		}
		channelMessage.ReplyToMessageID = HandlerUpdate.ChannelPost.MessageID
	} else {
		channelMessage.Text = user.UniqueId
		channelMessage.ReplyToMessageID = HandlerUpdate.ChannelPost.MessageID
	}
}

func channelHelp() {
	channelMessage.Text = utils.GetHelpTemplate(HandlerUpdate.ChannelPost.Chat.Title)
}

func ChannelHandler() {
	switch HandlerUpdate.ChannelPost.Text {
	case "!get":
		channelIdGenerator()
	case "!help":
		channelHelp()
	}
	channelMessage.ChatID = HandlerUpdate.ChannelPost.Chat.ID
	_, _ = HandlerBot.Send(channelMessage)
	channelMessage.Text = ""
	channelMessage.ReplyToMessageID = -1
}
