package entities

import (
	"strconv"

	"github.com/minhajuddinkhan/webrung/store/models"
)

//Game Game entity
type Game struct {
	GameID        string `json:"game_id"`
	PlayersJoined int    `json:"players_joined"`
}

//Marshal Marshal
func (game *Game) Marshal(g *models.Game) {

	gameID := strconv.FormatUint(uint64(g.Model.ID), 10)
	game.GameID = gameID
	game.PlayersJoined = g.PlayersJoined
}
