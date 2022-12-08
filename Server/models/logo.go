package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Logo struct {
	gorm.Model
	ImagePath string `json:"Imagepath" gorm:"type:varchar(300)"`
	Name      string `json:"logo_name" gorm:"type:varchar(300)"`
}

func GetRandomLogo() error {
	var logos []Logo
	db.Limit(4).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RAND()", Vars: []interface{}{}, WithoutParentheses: true},
	}).Find(&logos)
	fmt.Println(logos)
	return nil
}
