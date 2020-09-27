package slack

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func SendSlackMessage(message string, channel string) {
	api := slack.New(os.Getenv("SLACK_BOT"))
	channelID, timestamp, err := api.PostMessage(
		channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func SendSlackPackage(channel string, msg string, packageName string) error {
	api := slack.New(os.Getenv("SLACK_BOT"))
	params := slack.FileUploadParameters{
		Title:    packageName,
		Filetype: "apk",
		File:     "cache/" + packageName + ".apk",
	}
	file, err := api.UploadFile(params)
	if err != nil {
		SendSlackErrorMessage(channel)
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.EditLink)
	SendSlackMessage("msg"+file.Permalink, channel)
	return nil
}

func SendSlackErrorMessage(channel string) {
	SendSlackMessage("There was an error in sending your apk, please kindly check Actions section of your repository", channel)
}
