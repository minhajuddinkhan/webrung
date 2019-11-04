package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Winner        Player
	PlayersJoined int
}

//GetID gets id in string
func (g *Game) GetID() string {
	return strconv.Itoa(int(g.Model.ID))

}
