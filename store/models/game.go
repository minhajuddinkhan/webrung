package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Winner   Player
	WinnerID int
}

//GetID gets id in string
func (g *Game) GetID() string {
	return strconv.Itoa(int(g.Model.ID))

}
func (g *Game) SetID(ID string) error {
	n, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}
	g.ID = uint(n)
	return nil
}
