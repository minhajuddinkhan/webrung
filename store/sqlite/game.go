package sqlite

import (
	"github.com/davecgh/go-spew/spew"
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

func (sqlite *Game) GetJoinableGames() ([]models.JoinableGame, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.
		Table("players_in_games").
		Select("game_id, COUNT(*), player_id").
		Group("game_id").
		Having("count(*) < 4").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gameID, playerID string
	var playersJoined int
	var joinableGames []models.JoinableGame
	for rows.Next() {
		err := rows.Scan(&gameID, &playersJoined, &playerID)
		if err != nil {
			return nil, err
		}
		joinableGames = append(joinableGames, models.JoinableGame{
			GameID:        gameID,
			PlayersJoined: playersJoined,
			PlayerID:      playerID,
		})
	}

	return joinableGames, nil

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

func (sqlite *Game) IsPlayerInGame(playerID string) (bool, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return false, err
	}
	defer db.Close()
	var playerInGame models.PlayersInGame
	err = db.Where("player_id = ?", playerID).First(&playerInGame).Error
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

func (sqlite *Game) GetGameByHost(hostID string) (*models.Game, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var game models.Game
	if err := db.Where("host_id = ?", hostID).First(&game).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &errors.ErrGameIDNotFound{}
		}
		return nil, err
	}
	return &game, nil
}

//GetGameByPlayer GetGameByPlayer
func (sqlite *Game) GetGameByPlayer(playerID string) (*models.Game, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Table("players_in_games").
		Select("games.id, games.host_id").
		Joins("LEFT JOIN games  ON games.id = players_in_games.game_id").
		Where("players_in_games.player_id = ?", playerID).
		Rows()

	if err != nil {
		spew.Dump(err)
		return nil, err
	}

	var gameID, hostID uint
	for rows.Next() {
		if err := rows.Scan(&gameID, &hostID); err != nil {
			return nil, err
		}
	}

	return &models.Game{
		Model: gorm.Model{
			ID: gameID,
		},
		HostID: hostID,
	}, nil
}
