package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	User   User `json:"-"`
	UserID uint `json:"userId"`
	Score  uint `json:"max_score" gorm:"type:INT"`
}
