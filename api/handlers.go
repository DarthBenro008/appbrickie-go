package api

import (
	"appbrickie/api/database"
	"appbrickie/bot/slack"
	"appbrickie/bot/telegram/handlers"
	"github.com/gofiber/fiber"
	"log"
	"os"
	"strings"
)

//Router
func HandlerRouter(app fiber.Router) {
	app.Get("/status", statusCheck)
	app.Get("/sendMessage", sendID)
	app.Post("/sendPackage", sendPackage)
}

//HandlerFunctions
func statusCheck(c *fiber.Ctx) {
	_ = c.JSON(&fiber.Map{
		"msg":     handlers.HandlerBot.Self.UserName + " is up and running!",
		"success": true,
	})
}

func sendID(c *fiber.Ctx) {
	id := c.Query("id")
	msg := c.Query("msg")
	chatId, resp := database.ServiceHelper.GetUserChatId(id)
	if !resp {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"id":      id,
			"err":     "Error fetching chatId",
		})
	}
	res, err := handlers.SendMessage(chatId, msg)
	_ = c.JSON(&fiber.Map{
		"success": res,
		"id":      id,
		"err":     err,
	})

}

func sendPackage(c *fiber.Ctx) {
	id := c.FormValue("id")
	if strings.HasPrefix(id, "00") {
		sendSlackPackage(c, id)
	} else {
		sendTelegramPackage(c, id)
	}
}

func sendTelegramPackage(c *fiber.Ctx, id string) {
	msg := c.FormValue("msg")
	file, err := c.FormFile("file")
	chatId, resp := database.ServiceHelper.GetUserChatId(id)
	if !resp {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Invalid ID",
		})
		return
	}

	if err != nil {
		handlers.SendErrorMessage(chatId)
		log.Println("Invalid Multipart")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if file.Size > 50000000 {
		handlers.SendErrorMessage(chatId)
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Exceeded Size Limit!",
		})
		return
	}
	filename := file.Filename
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err = os.Mkdir("cache", 0700)
		if err != nil {
			log.Println(err)
		}
	}
	err = c.SaveFile(file, "cache/"+file.Filename)
	if err != nil {
		handlers.SendErrorMessage(chatId)
		log.Println("Error Saving File")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	response, err := handlers.SendPackage(chatId, msg, "cache/"+file.Filename)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"file":    filename,
			"error":   err.Error(),
		})
		return
	}
	_ = c.JSON(&fiber.Map{
		"success": response,
		"file":    filename,
	})
}

func sendSlackPackage(c *fiber.Ctx, id string) {
	msg := c.FormValue("msg")
	file, err := c.FormFile("file")
	channelId, resp := database.ServiceHelper.GetSlackUserChannelId(id)
	if !resp {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Invalid ID",
		})
		return
	}

	if err != nil {
		slack.SendSlackErrorMessage(channelId)
		log.Println("Invalid Multipart")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if file.Size > 50000000 {
		slack.SendSlackErrorMessage(channelId)
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Exceeded Size Limit!",
		})
		return
	}
	filename := file.Filename
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err = os.Mkdir("cache", 0700)
		if err != nil {
			log.Println(err)
		}
	}
	err = c.SaveFile(file, "cache/"+file.Filename)
	if err != nil {
		slack.SendSlackErrorMessage(channelId)
		log.Println("Error Saving File")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	err = slack.SendSlackPackage(channelId, msg, file.Filename)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"file":    filename,
			"error":   err.Error(),
		})
		return
	}
	_ = c.JSON(&fiber.Map{
		"success": true,
		"file":    filename,
	})
}
