package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect_db() {
	d, err := gorm.Open(mysql.Open("root:qwer1234@tcp(127.0.0.1:3306)/Guess_The_Logo?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
