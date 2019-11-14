package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

//PlayersInGame PlayersInGame
type PlayersInGame struct {
	gorm.Model
	GameID   uint
	Game     Game
	PlayerID uint
	Player   Player
}

//GetPlayerID GetPlayerID
func (pg *PlayersInGame) GetPlayerID() string {
	return strconv.Itoa(int(pg.PlayerID))
}

//JoinableGame JoinableGame
//this isnt not a table, hence should not be migrated.
type JoinableGame struct {
	GameID        string
	PlayersJoined int
	PlayerID      string
}
