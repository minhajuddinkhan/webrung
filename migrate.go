package webrung

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Migrate Migrate
func Migrate(dialect, connStr string) error {

	db, err := gorm.Open(dialect, connStr)
	if err != nil {
		return err
	}
	models := []interface{}{
		models.Card{},
		models.Game{},
		models.Player{},
		models.PlayersInGame{},
	}
	if err := db.AutoMigrate(models...).Error; err != nil {
		return err
	}
	return nil
}
