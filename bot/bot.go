package bot

import (
	"appbrickie/bot/handlers"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var (
	bot *tgbotapi.BotAPI
)

func InitialiseBot() {
	fmt.Println("Bot Starting ")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
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
