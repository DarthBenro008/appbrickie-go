package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);"`
	Username string `gorm:"type:varchar(80);"`
	ChatId   int64  `gorm:"type:numeric(40);primary_key"`
	UniqueId string `gorm:"type:varchar(100);unique"`
	Int      int    `gorm:"AUTO_INCREMENT"`
}
