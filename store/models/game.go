package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Winner   Player
	WinnerID uint
	Host     Player
	HostID   uint
}

//GetID gets id in string
func (g *Game) GetID() string {
	return idToString(g.Model.ID)

}

//SetID sets id in uint
func (g *Game) SetID(ID string) error {
	n, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}
	g.ID = uint(n)
	return nil
}

//GetHostID gets host id in string
func (g *Game) GetHostID() string {
	return idToString(g.HostID)
}

func idToString(id uint) string {
	return strconv.Itoa(int(id))
}
