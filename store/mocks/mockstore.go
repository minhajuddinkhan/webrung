package mocks

import (
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Store mock store
type GameStore struct {
	connErr bool
	game    models.Game
	player  models.Player
}

type PlayerStore struct {
	connErr bool
	game    models.Game
	player  models.Player
}

//NewStore creates a mock store for testing
func NewGameStore(connErr bool) (*GameStore, error) {
	return &GameStore{connErr: connErr}, nil
}

func NewPlayerStore(connErr bool) (*PlayerStore, error) {
	return &PlayerStore{connErr: connErr}, nil
}
