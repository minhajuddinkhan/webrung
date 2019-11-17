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
	ImageURL string
}

func (p *Player) SetID(ID uint) error {
	p.Model.ID = ID
	return nil
}
