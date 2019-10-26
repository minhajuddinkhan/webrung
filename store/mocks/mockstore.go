package mocks

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Store mock store
type Store struct {
	connErr bool
}

//NewStore creates a mock store for testing
func NewStore(connErr bool) (*Store, error) {
	return &Store{connErr: connErr}, nil
}

//CreateGame creates a mock game with id 123
func (ms *Store) CreateGame() (string, error) {
	if ms.connErr {
		return "", &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}
	return "123", nil
}

//GetGame GetGame
func (ms *Store) GetGame(gameID string) (*models.Game, error) {
	if ms.connErr {
		return nil, &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}
	if gameID == "69" {
		return &models.Game{
			Model: gorm.Model{
				ID: 69,
			},
			Players: []models.Player{{}},
		}, nil
	}

	return nil, &errors.ErrGameIDNotFound{GameID: gameID}

}

//Migrate Migrate
func (ms *Store) Migrate() error {
	return nil
}
