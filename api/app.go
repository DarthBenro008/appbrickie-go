package api

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"log"
)

func InitialiseApi() {
	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	HandlerRouter(api)
	log.Print("API Running")
	_ = app.Listen(3000)
}
