package game

import (
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

//GameManager GameManager
type GameManager interface {

	//CreateGame creates a game
	CreateGame(player *entities.Player) (game *entities.Game, err error)

	//GetGame gets a game
	GetGame(gameID string) (*entities.Game, error)

	//JoinGame joins a game
	JoinGame(player *entities.Player, game *entities.Game) error

	//StartGame starts a game
	StartGame(gameID string, startBy *entities.Player) ([]entities.Player, error)
}

type gameManager struct {
	store    store.Game
	ioclient iorpc.Client
}

//NewGameManager returns a new game manager.
func NewGameManager(s store.Game, c iorpc.Client) GameManager {
	return &gameManager{store: s, ioclient: c}
}
