package sqlite

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

type game struct {
	dialect string
	connStr string
}

func NewGameStore(connStr string) *game {
	return &game{dialect: "sqlite3", connStr: connStr}
}

//CreateGame CreateGame
func (sqlite *game) CreateGame(createdBy *models.Player) (string, error) {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	game := models.Game{
		Winner: models.Player{},
	}

	if err := db.Create(&game).Error; err != nil {
		return "", err
	}
	return game.GetID(), nil
}

//GetGame GetGame
func (sqlite *game) GetGame(gameID string) (*models.Game, error) {
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

func (sqlite *game) IncrementPlayersJoined(gameID string) error {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	n, _ := strconv.Atoi(gameID)
	fmt.Println()
	err = db.Exec("UPDATE games as g SET players_joined = g.players_joined + 1 WHERE id = ?", uint(n)).Error
	if err != nil {
		fmt.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &errors.ErrGameIDNotFound{GameID: gameID}
		}
		return err
	}
	return nil

}

//JoinGame JoinGame
func (sqlite *game) JoinGame(gameplay *models.PlayersInGame) error {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Create(gameplay).Error
}

func (sqlite *game) IsPlayerInGame(gameID string, playerID string) (bool, error) {

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

func (sqlite *game) GetPlayersInGame(gameID string) ([]models.PlayersInGame, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var players []models.PlayersInGame
	if err := db.Where("game_id = ?", gameID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil

}

func (sqlite *game) UpdateGame(gameID string, game *models.Game) error {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Where("id = ?", gameID).Update(game).Error

}
