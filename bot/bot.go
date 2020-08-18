package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitialiseBot() {
	fmt.Println("We Good ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env File")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error Initialising bot due to %s", err.Error())
	}
	bot.Debug = true
	log.Printf("%s Bot is up and running", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "sayhi":
			messageHandler(update, bot)
		default:
			msg.Text = "I do not know that command !"
		}
	}
}
