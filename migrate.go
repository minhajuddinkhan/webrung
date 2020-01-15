package webrung

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

func Migrate(dialect, connString string) {

	db, err := gorm.Open(dialect, connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(models.Models()...).Error; err != nil {
		log.Fatal(err)
	}
}
