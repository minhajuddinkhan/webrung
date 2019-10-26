package game

import (
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/store"
)

//GameManager GameManager
type GameManager interface {
	CreateGame() (game *entities.Game, err error)
	GetGame(gameID string) (*entities.Game, error)
}

type gameManager struct {
	store store.Store
}

//NewGameManager returns a new game manager.
func NewGameManager(s store.Store) GameManager {
	return &gameManager{store: s}
}

func (g *gameManager) CreateGame() (*entities.Game, error) {
	gameID, err := g.store.CreateGame()
	if err != nil {
		return nil, err
	}
	return &entities.Game{GameID: gameID, PlayersJoined: 0}, nil

}

func (g *gameManager) GetGame(gameID string) (*entities.Game, error) {
	game, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err

	}
	gameEntity := entities.Game{}
	gameEntity.Marshal(game)
	return &gameEntity, nil

}
