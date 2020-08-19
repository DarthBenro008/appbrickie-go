package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	bot *tgbotapi.BotAPI
)

type Bot interface {
	sendMessage() bool
}

type SendMessage struct {
	Id   int64
	Cmsg string
}

func NewBotInterface(inter Bot) {
	inter.sendMessage()
}

func InitialiseBot() {
	fmt.Println("We Good ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env File")
	}
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error Initialising bot due to %s", err.Error())
	}
	bot.Debug = true
	log.Printf("%s Bot is up and running", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		//TODO: Add Message and Channel Handlers
		if update.Message == nil {
			if update.ChannelPost != nil {
				if update.ChannelPost.Text == "!get" {
					msg := tgbotapi.NewMessage(update.ChannelPost.Chat.ID, strconv.Itoa(int(update.ChannelPost.Chat.ID)))
					msg.ReplyToMessageID = update.ChannelPost.MessageID
					_, _ = bot.Send(msg)
				}
			}
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "getID":
			messageHandler(update, bot)

		default:
			msg.Text = "I do not know that command !"
		}
	}
}
func (mdet SendMessage) sendMessage() bool {
	msg := tgbotapi.NewMessage(mdet.Id, mdet.Cmsg)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
