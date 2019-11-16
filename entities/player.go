package entities

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Player Player
type Player struct {
	ID     string `json:"id"`
	GameID string `json:"game_id"`
	Name   string `json:"name"`
	Cards  []Card `json:"cards"`
	InGame bool   `json:"in_game"`
	IsHost bool   `json:"is_host"`
}

func (p *Player) ToModel() (*models.Player, error) {

	id, err := strconv.Atoi(p.ID)
	if err != nil {
		return nil, err
	}

	return &models.Player{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name: p.Name,
	}, nil
}
