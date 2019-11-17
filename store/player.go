package store

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/store/models"
	"github.com/minhajuddinkhan/webrung/store/sqlite"
)

type Player interface {

	//CreatePlayer creates a new player
	CreatePlayer(player *models.Player) (playerID uint, err error)

	//GetPlayer gets a player
	GetPlayer(playerID uint) (*models.Player, error)

	//GetPlayerByName gets a player by name
	GetPlayerByName(name string) (*models.Player, error)
}

// NewPlayerStore NewPlayerStore
func NewPlayerStore(dialect, connString string) (Player, error) {

	if dialect == "sqlite3" {
		return sqlite.NewPlayerStore(connString), nil
	}

	return nil, fmt.Errorf("invalid dialect provided")
}
