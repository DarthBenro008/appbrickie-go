package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(80);unique_index"`
	ChatId   int64  `gorm:"type:numeric(40);primary_key"`
	UniqueId string `gorm:"type:varchar(100);unique"`
	Int      int    `gorm:"AUTO_INCREMENT"`
}
