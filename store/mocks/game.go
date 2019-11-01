package mocks

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreateGame creates a mock game with id 123
func (ms *Store) CreateGame() (string, error) {
	if ms.connErr {
		return "", &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}
	dummyID := 123
	ms.game = models.Game{PlayersJoined: 0, Model: gorm.Model{ID: uint(dummyID)}}
	return string(dummyID), nil
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
			PlayersJoined: 0,
		}, nil
	}
	return nil, &errors.ErrGameIDNotFound{GameID: gameID}
}
