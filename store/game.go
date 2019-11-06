package store

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/minhajuddinkhan/webrung/store/models"
	"github.com/minhajuddinkhan/webrung/store/sqlite"
)

//Game game store
type Game interface {
	//CreateGame creates a new game
	CreateGame(createdBy *models.Player) (gameID string, err error)

	//GetGame gets a game by id
	GetGame(gameID string) (*models.Game, error)

	//IsPlayerInGame returns if player is joined in a game
	IsPlayerInGame(gameID string, playerID string) (bool, error)

	//IncrementPlayersJoined IncrementPlayersJoined
	IncrementPlayersJoined(gameID string) error

	JoinGame(gameplay *models.PlayersInGame) error
}

//NewGameStore NewGameStore
func NewGameStore(dialect, connStr string) (Game, error) {
	if dialect == "sqlite3" {
		return sqlite.NewGameStore(connStr), nil
	}

	if dialect == "mock" {
		return mocks.NewGameStore(false)
	}

	return nil, fmt.Errorf("invalid dialect")
}
