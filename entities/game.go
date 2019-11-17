package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Game Game entity
type Game struct {
	GameID        uint `json:"game_id"`
	PlayersJoined int  `json:"players_joined,omitempty"`
	HostID        uint `json:"host_id"`
}

//ToModel ToModel
func (game *Game) ToModel() (*models.Game, error) {

	return &models.Game{
		Model: gorm.Model{
			ID: game.GameID,
		},
		HostID: game.HostID,
	}, nil
}
