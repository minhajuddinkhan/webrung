package store

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/store/models"
	"github.com/minhajuddinkhan/webrung/store/sqlite"
)

//Store Store
type Store interface {
	CreateGame() (gameID string, err error)
	GetGame(gameID string) (*models.Game, error)

	//Migrate migrates all tables
	Migrate() error
}

//NewRungStore creates a new rung store
func NewRungStore(dialect, connStr string) (Store, error) {

	if dialect == "sqlite3" {
		return sqlite.NewStore(connStr)
	}

	return nil, fmt.Errorf("%s dialect not supported", dialect)
}
