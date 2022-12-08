package models

import (
	"github.com/J-Nokwal/Guess_The_Logo/Server/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	configs.Connect()
	db = configs.GetDB()
	db.AutoMigrate(&User{}, &User{}, &Logo{}, &Report{}, &Game{})
}
