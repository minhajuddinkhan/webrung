package sqlite

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

type playerStore struct {
	connStr string
	dialect string
}

//CreatePlayer CreatePlayer
func (ps *playerStore) CreatePlayer(pl *models.Player) (string, error) {

	db, err := gorm.Open(ps.dialect, ps.connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()

	if err := db.Create(pl).Error; err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(pl.Model.ID), 10), nil
}

func (ps *playerStore) GetPlayer(playerID string) (*models.Player, error) {
	db, err := gorm.Open(ps.dialect, ps.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	player := models.Player{}
	if err := db.Where("id = ?", playerID).First(&player).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &errors.ErrPlayerNotFound{}
		}
		//TODO:: add player not found error
		return nil, err
	}
	return &player, nil

}

//GetPlayerByName GetPlayerByName
func (ps *playerStore) GetPlayerByName(name string) (*models.Player, error) {
	db, err := gorm.Open(ps.dialect, ps.connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	player := models.Player{}
	if err := db.Where("name = ?", name).First(&player).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &errors.ErrPlayerNotFound{}
		}
		return nil, err
	}
	return &player, nil

}
