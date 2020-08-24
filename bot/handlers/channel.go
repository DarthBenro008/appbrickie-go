package handlers

import (
	"appbrickie/bot/utils"
	"strconv"
)

func channelGreet() {
	channelMessage.Text = strconv.Itoa(int(HandlerUpdate.ChannelPost.Chat.ID))
	channelMessage.ReplyToMessageID = HandlerUpdate.ChannelPost.MessageID
}

func channelHelp() {
	channelMessage.Text = utils.GetHelpTemplate(HandlerUpdate.ChannelPost.Chat.Title)
}

func ChannelHandler() {
	switch HandlerUpdate.ChannelPost.Text {
	case "!get":
		channelGreet()
	case "!help":
		channelHelp()
	}
	channelMessage.ChatID = HandlerUpdate.ChannelPost.Chat.ID
	_, _ = HandlerBot.Send(channelMessage)
}
