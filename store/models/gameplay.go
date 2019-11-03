package models

import "github.com/jinzhu/gorm"

//PlayersInGame PlayersInGame
type PlayersInGame struct {
	gorm.Model
	gameID   uint
	Game     Game
	PlayerID uint
	Player   Player
}
