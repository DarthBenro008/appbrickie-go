package api

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"log"
)

func InitialiseApi() {
	app := fiber.New()
	app.Use(cors.New())
	app.Settings.BodyLimit = 52428800
	api := app.Group("/api")
	app.Use(middleware.Logger())
	HandlerRouter(api)
	log.Print("API is up and Running")
	_ = app.Listen(3000)
}
