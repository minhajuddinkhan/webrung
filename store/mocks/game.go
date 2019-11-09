package mocks

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreateGame creates a mock game with id 123
func (ms *GameStore) CreateGame(pl *models.Player) (string, error) {
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
		g := &models.Game{}
		g.SetID(gameID)
		return g, nil
	}
	return nil, &errors.ErrGameIDNotFound{}
}

//JoinGame JoinGame
func (ms *GameStore) JoinGame(gameplay *models.PlayersInGame) error {

	if len(ms.playersInGame) == 4 {
		return fmt.Errorf("4 players already")
	}
	ms.playersInGame = append(ms.playersInGame, *gameplay)
	return nil
}

//IsPlayerInGame IsPlayerInGame
func (ms *GameStore) IsPlayerInGame(gameID string, playerID string) (bool, error) {

	//TODO:: devise strategy.
	return false, nil
}

func (ms *GameStore) GetPlayersInGame(gameID string) (players []models.PlayersInGame, err error) {

	return ms.playersInGame, nil

}

func (ms *GameStore) StartGame(gameID string) ([]entities.Player, error) {
	return nil, nil
}
