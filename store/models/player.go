package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

//Player player in a rung game
type Player struct {
	gorm.Model
	Name     string
	Cards    []Card
	InGame   uint
	HandsWon int
}

func (p *Player) SetID(ID string) error {
	id, err := strconv.Atoi(ID)
	p.Model.ID = uint(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) GetID() string {
	return strconv.Itoa(int(p.Model.ID))
}
