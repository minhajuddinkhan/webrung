package store

import (
	"github.com/minhajuddinkhan/webrung/store/models"
)

type Player interface {

	//CreatePlayer creates a new player
	CreatePlayer(player *models.Player) (playerID string, err error)

	//GetPlayer gets a player
	GetPlayer(playerID string) (*models.Player, error)

	//GetPlayerByName gets a player by name
	GetPlayerByName(name string) (*models.Player, error)
}
