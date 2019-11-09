package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

type Game struct {
	dialect string
	connStr string
}

//NewGameStore creates new game store
func NewGameStore(connStr string) *Game {
	return &Game{dialect: "sqlite3", connStr: connStr}
}

//CreateGame CreateGame
func (sqlite *Game) CreateGame(createdBy *models.Player) (string, error) {
	//TODO:: use this createdby to store host of the game
	//only the host of the game should be able to start the game.

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	game := models.Game{
		HostID: createdBy.ID,
		Winner: models.Player{},
	}

	if err := db.Create(&game).Error; err != nil {
		return "", err
	}
	return game.GetID(), nil
}

//GetGame GetGame
func (sqlite *Game) GetGame(gameID string) (*models.Game, error) {
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

//JoinGame JoinGame
func (sqlite *Game) JoinGame(gameplay *models.PlayersInGame) error {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Create(gameplay).Error
}

func (sqlite *Game) IsPlayerInGame(gameID string, playerID string) (bool, error) {

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

func (sqlite *Game) GetPlayersInGame(gameID string) ([]models.PlayersInGame, error) {

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
