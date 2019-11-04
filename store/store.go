package store

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/minhajuddinkhan/webrung/store/models"
	"github.com/minhajuddinkhan/webrung/store/sqlite"
)

//Store Store
type Store interface {
	//CreateGame creates a new game
	CreateGame(createdBy *models.Player) (gameID string, err error)

	//GetGame gets a game by id
	GetGame(gameID string) (*models.Game, error)

	//IsPlayerInGame returns if player is joined in a game
	IsPlayerInGame(gameID string, playerID string) (bool, error)

	//IncrementPlayersJoined IncrementPlayersJoined
	IncrementPlayersJoined(gameID string) error

	JoinGame(gameplay *models.PlayersInGame) error

	//CreatePlayer creates a new player
	CreatePlayer(player *models.Player) (playerID string, err error)

	//GetPlayer gets a player
	GetPlayer(playerID string) (*models.Player, error)

	//GetPlayerByName gets a player by name
	GetPlayerByName(name string) (*models.Player, error)

	// IncrementGamePlayer() (*models.Game, error)
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
