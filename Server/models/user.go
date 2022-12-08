package models

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(300)"`
	MaxScore uint   `json:"max_score" gorm:"type:INT ;default:0"`
}

func (user *User) CreateUser() (*User, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("error while creating User")
	}
	return user, nil
}

func (user *User) SetNewScore() error {
	var usertemp User
	if err := db.Where("id=?", user.ID).Find(&usertemp); err != nil {
		return fmt.Errorf("error User Not Found")
	}
	usertemp.MaxScore = uint(math.Max(float64(user.MaxScore), float64(user.MaxScore)))
	if err := db.Model(&usertemp).Update("max_score", usertemp.MaxScore); err != nil {
		return fmt.Errorf("error While updating user")
	}
	return nil
}
