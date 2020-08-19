package api

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"log"
)

func InitialiseApi() {
	app := fiber.New()

	//app.Use(cors.New())
	api := app.Group("/api")
	app.Use(middleware.Logger())
	HandlerRouter(api)
	log.Print("API Running")
	_ = app.Listen(3000)
}
