package db

import (
	"appbrickie/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

func InitialiseDatabase() {
	dbargs := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", dbargs)
	if err != nil {
		log.Fatal("Error Connecting to database", err.Error())
	}
	log.Println("Connected to Database")
	db.AutoMigrate(&models.User{})
	user := models.User{Username: "Test", ChatId: 111111, UniqueId: 222222}
	check := db.Create(user)
	log.Println(check)
	db.Create(&user).Update("CreatedAt", time.Now())
	defer db.Close()
}
