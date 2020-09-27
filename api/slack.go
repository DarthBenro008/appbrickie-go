package api

import (
	"appbrickie/api/database"
	"appbrickie/api/database/models"
	"appbrickie/bot/slack"
	"github.com/gofiber/fiber"
	"log"
)

func SlackRouter(app fiber.Router) {
	app.Post("/getuid", getID)

}

func getID(c *fiber.Ctx) {
	incomingData := new(models.Slack)
	if err := c.BodyParser(incomingData); err != nil {
		log.Fatal(err)
	}
	user, resp := database.ServiceHelper.FetchSlackUser(incomingData.ChannelID)
	if !resp {
		uid, success := database.ServiceHelper.CreateSlackUser(incomingData.TeamName, incomingData.ChannelName, incomingData.ChannelID, incomingData.TeamID)
		if success {
			slack.SendSlackMessage(uid, incomingData.ChannelID)
		} else {
			slack.SendSlackMessage("Error generating your uid, please try again later", incomingData.ChannelID)
		}
	} else {
		slack.SendSlackMessage(user.UniqueId, incomingData.ChannelID)
	}
}
