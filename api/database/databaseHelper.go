package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ServiceHelper Service
)
func InitialiseDatabase(service Service){
	ServiceHelper = service
}
