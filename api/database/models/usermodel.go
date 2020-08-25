package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);"`
	Username string `gorm:"type:varchar(80);"`
	ChatId   string `gorm:"type:varchar(80);primary_key"`
	UniqueId string `gorm:"type:varchar(100);unique"`
	Channel  bool   `gorm:"type:boolean"`
	Int      int    `gorm:"AUTO_INCREMENT"`
}
