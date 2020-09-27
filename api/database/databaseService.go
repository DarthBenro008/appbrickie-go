package database

import (
	"appbrickie/api/database/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"log"
	"strconv"
	"time"
)

type Service interface {
	CreateUser(chatId int64, username string, name string, isChannel bool) (string, bool)
	GetUserChatId(uid string) (int64, bool)
	FetchUser(chatId int64) (models.User, bool)
	CreateSlackUser(teamName string, channelName string, channelId string, teamId string) (string, bool)
	GetSlackUserChannelId(uid string) (string, bool)
	FetchSlackUser(channelId string) (models.Slack, bool)
}

type databaseService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &databaseService{db: db}
}

func (s *databaseService) CreateSlackUser(teamName string, channelName string, channelId string, teamId string) (string, bool) {
	guid := "00" + xid.New().String()
	user := models.Slack{TeamName: teamName, ChannelName: channelName, ChannelID: channelId, TeamID: teamId, UniqueId: guid}
	err := s.db.Create(&user).Update("CreatedAt", time.Now()).Error
	if err != nil {
		return err.Error(), false
	}
	return "", true
}

func (s *databaseService) GetSlackUserChannelId(uid string) (string, bool) {
	var user models.Slack
	result := s.db.Where("unique_id = ?", uid).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return "", false
	}
	return user.ChannelID, true
}

func (s *databaseService) FetchSlackUser(channelId string) (models.Slack, bool) {
	var user models.Slack
	result := s.db.Where("channel_id = ?", channelId).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return user, false
	}
	return user, true

}

func (s *databaseService) CreateUser(chatId int64, username string, name string, isChannel bool) (string, bool) {
	guid := xid.New()
	rid := strconv.Itoa(int(chatId))
	if len(username) == 0 {
		username = "No Username"
	}
	user := models.User{Username: username, ChatId: rid, UniqueId: guid.String(), Name: name, Channel: isChannel}
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
	rid, _ := strconv.ParseInt(user.ChatId, 10, 64)
	return rid, true
}

func (s *databaseService) FetchUser(chatId int64) (models.User, bool) {
	var user models.User
	result := s.db.Where("chat_id = ?", chatId).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return user, false
	}
	return user, true
}
