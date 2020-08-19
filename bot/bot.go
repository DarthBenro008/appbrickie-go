package bot

import (
	"appbrickie/bot/handlers"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	handlers.HandlerBot = bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		handlers.HandlerUpdate = update
		if update.ChannelPost != nil {
			handlers.ChannelHandler()
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		handlers.ChatHandler()
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
