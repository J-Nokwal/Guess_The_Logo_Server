package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	User   User `json:"-"`
	UserID uint `json:"userId"`
	Score  uint `json:"max_score" gorm:"type:INT"`
}

func (game Game) save() error {
	if err := db.Create(game).Error; err != nil {
		return fmt.Errorf("Error In Inserting Game")
	}
	return nil
}

func GetTopScorers() ([]Game, error) {
	topScores := make([]Game, 10)
	if err := db.Model(&Game{}).Order("score desc").Limit(10).Find(&topScores).Error; err != nil {
		return nil, fmt.Errorf("Error while Quering top Scores")
	}
	return topScores, nil
}
