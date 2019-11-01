package mocks

import (
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Store mock store
type Store struct {
	connErr bool
	game    models.Game
	player  models.Player
}

//NewStore creates a mock store for testing
func NewStore(connErr bool) (*Store, error) {
	return &Store{connErr: connErr}, nil
}

//Migrate Migrate
func (ms *Store) Migrate() error {
	return nil
}
