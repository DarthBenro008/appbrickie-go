package api

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"log"
	"os"
)

func InitialiseApi() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Welcome to App Brickie Api!")
	})
	app.Settings.BodyLimit = 52428800
	api := app.Group("/api")
	//app.Use(middleware.Logger())
	HandlerRouter(api)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Print("API is up and Running on Port " + port)
	_ = app.Listen(port)

}
