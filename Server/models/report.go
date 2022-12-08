package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	User        User   `json:"-"`
	UserID      uint   `json:"userId" `
	Logo        Logo   `json:"-"`
	LogoID      uint   `json:"logoId" `
	Description string `json:"description"`
}
