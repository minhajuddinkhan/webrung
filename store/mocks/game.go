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
func (ms *GameStore) CreateGame(createdBy *models.Player) (string, error) {
	if ms.connErr {
		return "", &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}

	dummyID := 1
	ms.game = models.Game{
		Model: gorm.Model{ID: uint(dummyID)},
	}
	return strconv.Itoa(dummyID), nil
}

//GetGame GetGame
func (ms *GameStore) GetGame(gameID string) (*models.Game, error) {
	if ms.connErr {
		return nil, &errors.ErrDBConnection{ConnectionString: "mock failed connection"}
	}
	if gameID == strconv.Itoa(int(ms.game.ID)) {
		return &ms.game, nil
	}
	return nil, &errors.ErrGameIDNotFound{}
}

//JoinGame JoinGame
func (ms *GameStore) JoinGame(gameplay *models.PlayersInGame) error {
	mockGamePlayInStore = append(mockGamePlayInStore, *gameplay)
	return nil
}

//IncrementPlayersJoined IncrementPlayersJoined
func (ms *GameStore) IncrementPlayersJoined(gameID string) error {
	return nil
}

//IsPlayerInGame IsPlayerInGame
func (ms *GameStore) IsPlayerInGame(gameID string, playerID string) (bool, error) {

	//TODO:: devise strategy.
	return false, nil
}

func (ms *GameStore) GetPlayersInGame(gameID string) (players []models.PlayersInGame, err error) {

	//TODO:: add 4 players
	return []models.PlayersInGame{}, nil
}
