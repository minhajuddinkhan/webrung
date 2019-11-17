package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Player Player
type Player struct {
	ID     uint   `json:"id"`
	GameID uint   `json:"game_id"`
	Name   string `json:"name"`
	Cards  []Card `json:"cards"`
	InGame bool   `json:"in_game"`
	IsHost bool   `json:"is_host"`
}

func (p *Player) ToModel() (*models.Player, error) {

	return &models.Player{
		Model: gorm.Model{
			ID: p.ID,
		},
		Name: p.Name,
	}, nil
}
