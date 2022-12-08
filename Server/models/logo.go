package models

import "gorm.io/gorm"

type Logo struct {
	gorm.Model
	ImagePath string `json:"Imagepath" gorm:"type:varchar(300)"`
	Name      string `json:"logo_name" gorm:"type:varchar(300)"`
}
