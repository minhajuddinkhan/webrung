package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store"
	"github.com/minhajuddinkhan/webrung/store/models"
)

func main() {

	players := []models.Player{
		models.Player{
			Name:     "East",
			HandsWon: 0,
			Cards:    []models.Card{},
		},
		models.Player{
			Name:     "West",
			HandsWon: 0,
			Cards:    []models.Card{},
		},
		models.Player{
			Name:     "North",
			HandsWon: 0,
			Cards:    []models.Card{},
		},
		models.Player{
			Name:     "South",
			HandsWon: 0,
			Cards:    []models.Card{},
		},
	}
	dialect := os.Getenv("DB_DIALECT")
	dbConnString := os.Getenv("DB_URL")
	store, err := store.NewRungStore(dialect, dbConnString)
	if err != nil {
		log.Fatal(err)
	}

	err = store.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	for _, player := range players {
		_, err := store.GetPlayerByName(player.Name)

		if err != nil {
			switch err.(type) {
			case *(errors.ErrPlayerNotFound):
				_, err := store.CreatePlayer(&player)
				if err != nil {
					log.Fatal(err)
				}
				logrus.Info("player seeded!", player.Name)

			default:
				log.Fatal(err)
			}
		}
	}

	logrus.Info("Database seeded!")
}
