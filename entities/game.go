package entities

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Game Game entity
type Game struct {
	GameID        string `json:"game_id"`
	PlayersJoined int    `json:"players_joined,omitempty"`
}

//Marshal Marshal
func (game *Game) Marshal(g *models.Game) {
	gameID := strconv.FormatUint(uint64(g.Model.ID), 10)
	game.GameID = gameID
	game.PlayersJoined = g.PlayersJoined
}

//ToModel ToModel
func (game *Game) ToModel() (*models.Game, error) {

	id, err := strconv.Atoi(game.GameID)
	if err != nil {
		return nil, err
	}

	return &models.Game{
		Model: gorm.Model{
			ID: uint(id),
		},
	}, nil
}
