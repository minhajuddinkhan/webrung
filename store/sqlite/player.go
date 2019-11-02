package sqlite

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreatePlayer CreatePlayer
func (sqlite *Store) CreatePlayer(pl *models.Player) (string, error) {

	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	if err := db.Create(pl).Error; err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(pl.Model.ID), 10), nil
}

func (sqlite *Store) GetPlayer(playerID string) (*models.Player, error) {
	db, err := gorm.Open(sqlite.dialect, sqlite.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	player := models.Player{}
	if err := db.Where("id = ?", playerID).First(&player).Error; err != nil {
		//TODO:: add player not found error
		return nil, err
	}
	return &player, nil

}
