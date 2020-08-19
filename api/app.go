package api

import (
	"appbrickie/bot"
	"github.com/gofiber/fiber"
	"log"
	"strconv"
)

func InitialiseApi() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		_ = c.JSON(&fiber.Map{
			"msg": "HelloWorld",
		})
	})
	app.Get("/api/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		realid, _ := strconv.ParseInt(id, 10, 64)
		newMessage := bot.SendMessage{Id: realid, Cmsg: "Hello World!"}
		bot.NewBotInterface(newMessage)
		_ = c.JSON(&fiber.Map{
			"Id": id,
		})
	})
	log.Print("API Running")
	_ = app.Listen(3000)
}
