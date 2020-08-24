package database

import (
	"appbrickie/api/database/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"log"
	"time"
)

type Service interface {
	CreateUser(chatId int64, username string, name string) (string, bool)
	GetUserChatId(uid string) (int64, bool)
	FetchUser(chatId int64) (models.User, bool)
}

type databaseService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &databaseService{db: db}
}

func (s *databaseService) CreateUser(chatId int64, username string, name string) (string, bool) {
	guid := xid.New()
	if len(username) == 0 {
		username = "No Username"
	}
	user := models.User{Username: username, ChatId: chatId, UniqueId: guid.String(), Name: name}
	err := s.db.Create(&user).Update("CreatedAt", time.Now()).Error
	if err != nil {
		log.Println(err.Error())
		return err.Error(), false
	}
	return guid.String(), true
}

func (s *databaseService) GetUserChatId(uid string) (int64, bool) {
	var user models.User
	result := s.db.Where("unique_id = ?", uid).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return 0, false
	}
	return user.ChatId, true
}

func (s *databaseService) FetchUser(chatId int64) (models.User, bool) {
	var user models.User
	result := s.db.Where("chat_id = ?", chatId).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return user, false
	}
	return user, true
}
