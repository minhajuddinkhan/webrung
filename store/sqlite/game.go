package sqlite

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreateGame CreateGame
func (sqlite *Store) CreateGame() (string, error) {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	game := models.Game{
		Players: []models.Player{},
		Winners: []models.Player{},
	}

	if err := db.Create(&game).Error; err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(game.Model.ID), 10), nil
}

//GetGame GetGame
func (sqlite *Store) GetGame(gameID string) (*models.Game, error) {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var game models.Game
	if err := db.Where("id = ?", gameID).First(&game).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &errors.ErrGameIDNotFound{GameID: gameID}
		}
		return nil, err
	}
	return &game, nil
}
