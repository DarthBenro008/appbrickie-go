package handlers

import (
	"strconv"
)

func channelGreet() {
	channelMessage.Text = strconv.Itoa(int(HandlerUpdate.ChannelPost.Chat.ID))
	channelMessage.ReplyToMessageID = HandlerUpdate.ChannelPost.MessageID
}

func ChannelHandler() {
	if HandlerUpdate.ChannelPost.Text == "!get" {
		channelGreet()
	}
	channelMessage.ChatID = HandlerUpdate.ChannelPost.Chat.ID
	_, _ = HandlerBot.Send(channelMessage)
}
