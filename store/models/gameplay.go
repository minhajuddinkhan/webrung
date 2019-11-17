package models

import (
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

//JoinableGame JoinableGame
//this isnt not a table, hence should not be migrated.
type JoinableGame struct {
	GameID        uint
	PlayersJoined int
	PlayerID      uint
}
