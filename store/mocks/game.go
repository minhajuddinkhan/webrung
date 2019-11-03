package mocks

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

var mockGamePlayInStore []models.PlayersInGame
var counter int

//CreateGame creates a mock game with id 123
func (ms *Store) CreateGame(createdBy *models.Player) (string, error) {
	if ms.connErr {
		return "", &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}

	dummyID := 1
	ms.game = models.Game{
		PlayersJoined: 1,
		Model:         gorm.Model{ID: uint(dummyID)},
	}
	return strconv.Itoa(dummyID), nil
}

//GetGame GetGame
func (ms *Store) GetGame(gameID string) (*models.Game, error) {
	if ms.connErr {
		return nil, &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}
	if gameID == strconv.Itoa(int(ms.game.ID)) {
		return &ms.game, nil
	}
	return nil, &errors.ErrGameIDNotFound{}
}

//JoinGame JoinGame
func (ms *Store) JoinGame(gameplay *models.PlayersInGame) error {
	mockGamePlayInStore = append(mockGamePlayInStore, *gameplay)
	return nil
}

//IncrementPlayersJoined IncrementPlayersJoined
func (ms *Store) IncrementPlayersJoined(gameID string) error {
	ms.game.PlayersJoined++
	return nil
}
