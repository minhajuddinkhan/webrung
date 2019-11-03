package sqlite

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreateGame CreateGame
func (sqlite *Store) CreateGame(createdBy *models.Player) (string, error) {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	game := models.Game{
		PlayersJoined: 1,
		Winner:        models.Player{},
	}

	if err := db.Create(&game).Error; err != nil {
		return "", err
	}
	err = sqlite.JoinGame(&models.PlayersInGame{
		Game:   game,
		Player: *createdBy,
	})
	if err != nil {
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

func (sqlite *Store) IncrementPlayersJoined(gameID string) error {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Where("id = ?", gameID).Update("SET players_joined = players_joind + 1").Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &errors.ErrGameIDNotFound{GameID: gameID}
		}
		return err
	}
	return nil

}

//JoinGame JoinGame
func (sqlite *Store) JoinGame(gameplay *models.PlayersInGame) error {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Create(gameplay).Error
}
