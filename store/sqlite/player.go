package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

type playerStore struct {
	connStr string
	dialect string
}

//NewSqlitePlayerStore NewSqlitePlayerStore
func NewPlayerStore(connStr string) *playerStore {
	return &playerStore{connStr: connStr, dialect: "sqlite3"}
}

//CreatePlayer CreatePlayer
func (ps *playerStore) CreatePlayer(pl *models.Player) (uint, error) {

	db, err := gorm.Open(ps.dialect, ps.connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	if err := db.Create(pl).Error; err != nil {
		return 0, err
	}

	return pl.Model.ID, nil
}

func (ps *playerStore) GetPlayer(playerID uint) (*models.Player, error) {
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
