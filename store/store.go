package store

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/minhajuddinkhan/webrung/store/models"
	"github.com/minhajuddinkhan/webrung/store/sqlite"
)

//Store Store
type Store interface {
	CreateGame() (gameID string, err error)
	GetGame(gameID string) (*models.Game, error)
	//CreatePlayer creates a new player
	CreatePlayer(player *models.Player) (playerID string, err error)

	GetPlayer(playerID string) (*models.Player, error)
	//Migrate migrates all tables
	Migrate() error
}

//NewRungStore creates a new rung store
func NewRungStore(dialect, connStr string) (Store, error) {

	if dialect == "sqlite3" {
		return sqlite.NewStore(connStr)
	}
	if dialect == "mock" {
		return mocks.NewStore(false)
	}

	return nil, fmt.Errorf("%s dialect not supported", dialect)
}
