package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(300)"`
	MaxScore uint   `json:"max_score" gorm:"type:INT"`
}
