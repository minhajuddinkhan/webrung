package game

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/store/models"
)

func (g *gameManager) JoinGame(player *entities.Player, game *entities.Game) error {

	gameModel, err := game.ToModel()
	if err != nil {
		return err
	}
	playerModel, err := player.ToModel()
	if err != nil {
		return err
	}

	currentGame, err := g.store.GetGame(game.GameID)
	if err != nil {
		return err
	}

	if currentGame.PlayersJoined == 4 {
		return fmt.Errorf("already four players joined")
	}

	present, err := g.store.IsPlayerInGame(game.GameID, player.ID)
	if err != nil {
		return err
	}
	if present {
		return fmt.Errorf("player already present in game. cannot join again")
	}

	gameplay := models.PlayersInGame{
		Game:   *gameModel,
		Player: *playerModel,
	}

	if err := g.store.JoinGame(&gameplay); err != nil {
		return err
	}

	return g.store.IncrementPlayersJoined(game.GameID)
}
