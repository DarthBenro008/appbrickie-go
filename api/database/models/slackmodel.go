package models

import "github.com/jinzhu/gorm"

type Slack struct {
	gorm.Model
	TeamName    string `gorm:"type:varchar(100);" json:"team_domain" xml:"team_domain" form:"team_domain"`
	ChannelName string `gorm:"type:varchar(80);" json:"channel_name" xml:"channel_name" form:"channel_name"`
	ChannelID   string `gorm:"type:varchar(80);primary_key" json:"channel_id" xml:"channel_id" form:"channel_id"`
	TeamID      string `gorm:"type:varchar(100);" json:"team_id" xml:"team_id" form:"team_id"`
	UniqueId    string `gorm:"type:varchar(100);unique"`
	Int         int    `gorm:"AUTO_INCREMENT"`
}
