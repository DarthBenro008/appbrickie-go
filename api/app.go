package api

import (
	"github.com/gofiber/fiber"
	"log"
)

func InitialiseApi() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		_ = c.JSON(&fiber.Map{
			"msg": "HelloWorld",
		})
	})
	log.Print("API Running")
	_ = app.Listen(3000)
}
