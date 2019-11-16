package store

import (
	"fmt"

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

	JoinGame(gameplay *models.PlayersInGame) error

	//GetPlayersInGame gets all players that are in the game
	GetPlayersInGame(gameID string) (players []models.PlayersInGame, err error)

	//GetJoinableGames gets joinable games. i.e players with less than all 4 players joined.
	GetJoinableGames() ([]models.JoinableGame, error)

	//GetGameByHost get game by host
	GetGameByHost(hostID string) (*models.Game, error)
}

//NewGameStore NewGameStore
func NewGameStore(dialect, connStr string) (Game, error) {
	if dialect == "sqlite3" {
		return sqlite.NewGameStore(connStr), nil
	}

	return nil, fmt.Errorf("invalid dialect")
}
