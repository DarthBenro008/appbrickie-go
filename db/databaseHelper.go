package db

import (
	"appbrickie/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/xid"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

func InitialiseDatabase() {
	dbargs := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	var err error
	db, err = gorm.Open("postgres", dbargs)
	if err != nil {
		log.Fatal("Error Connecting to database", err.Error())
	}
	db.AutoMigrate(&models.User{})
	db.Close()
}

func CreateUser(chatid int64, username string) (string, bool) {
	guid := xid.New()
	user := models.User{Username: username, ChatId: chatid, UniqueId: guid.String()}
	err := db.Create(&user).Update("CreatedAt", time.Now()).Error
	if err != nil {
		log.Println(err.Error())
		return err.Error(), false
	}
	return guid.String(), true
}

func GetUserChatId(uid string) (int64, bool) {
	var user models.User
	result := db.Where("unique_id = ?", uid).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return 0, false
	}
	return user.ChatId, true
}
