package models

import (
	"github.com/jinzhu/gorm"
)

//Player player in a rung game
type Player struct {
	gorm.Model
	Name     string
	Cards    []Card
	HandsWon int
}
