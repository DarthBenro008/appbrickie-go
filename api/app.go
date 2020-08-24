package api

import (
	"appbrickie/api/database"
	"appbrickie/api/database/models"
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
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

	dbargs := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", dbargs)
	if err != nil {
		log.Fatal("Error Connecting to database", err.Error())
	}
	log.Println("Connected to database")

	//Migrate Tables
	db.AutoMigrate(&models.User{})
	defer db.Close()

	serviceHelper := database.NewService(db)
	database.InitialiseDatabase(serviceHelper)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Print("API is up and Running on Port " + port)
	_ = app.Listen(port)

}
