package sqlite

import (
	"fmt"

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
	return game.GetID(), nil
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

	err = db.Exec("UPDATE games SET players_joined = players_joined + 1 WHERE id = ?", gameID).Error
	// err = db.Model(&models.Game{}).Where("id = ?", gameID).Update("").Error
	if err != nil {
		fmt.Println("ERRRR", err)
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

func (sqlite *Store) IsPlayerInGame(gameID string, playerID string) (bool, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return false, err
	}
	defer db.Close()
	var playerInGame models.PlayersInGame
	err = db.Where("game_id = ? AND player_id = ?", gameID, playerID).First(&playerInGame).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
