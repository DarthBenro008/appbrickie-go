package slack

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/slack-go/slack"
	"log"
	"os"
)

type data struct {
	ChannelID string `json:"channel_id" xml:"channel_id" form:"channel_id"`
}

func SlackRouter(app fiber.Router) {
	app.Post("/hello", getID)

}
func getID(c *fiber.Ctx) {

}

func sendPackage(c *fiber.Ctx) {
	incomingData := new(data)
	if err := c.BodyParser(incomingData); err != nil {
		log.Fatal(err)
	}
	log.Print("Hello ", incomingData.ChannelID)
	api := slack.New(os.Getenv("SLACK_BOT"))

	params := slack.FileUploadParameters{
		Title:    "Example File",
		Filetype: "apk",
		File:     "lol.apk",
		//Content:  "Nan Nan Nan Nan Nan Nan Nan Nan Batman",
	}

	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.EditLink)

	channelID, timestamp, err := api.PostMessage(
		incomingData.ChannelID,
		slack.MsgOptionText(file.Permalink, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
