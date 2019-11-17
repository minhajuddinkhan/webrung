package game

import (
	"fmt"
	"strconv"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/store/models"
)

func (g *gameManager) CreateGame(pl *entities.Player) (*entities.Game, error) {

	_, err := g.store.GetGameByHost(pl.ID)
	var gameAlreadyHosted bool
	if err != nil {
		switch err.(type) {
		case (*errors.ErrGameIDNotFound):
			gameAlreadyHosted = false

		default:
			return nil, err
		}

	}
	if gameAlreadyHosted || err == nil {
		return nil, &errors.ErrGameAlreadyHosted{
			Err: fmt.Errorf("game already hosted by player %s", pl.ID),
		}
	}

	isJoinedInAnotherGame, err := g.store.IsPlayerInGame(pl.ID)
	if err != nil {
		return nil, err
	}

	if isJoinedInAnotherGame {
		return nil, &errors.ErrPlayerAlreadyJoinedInAnotherGame{
			Err: fmt.Errorf("game already joined by player %s", pl.ID),
		}
	}

	player := models.Player{Name: pl.Name}
	if err := player.SetID(pl.ID); err != nil {
		return nil, err
	}

	gameID, err := g.store.CreateGame(&player)
	if err != nil {
		return nil, err
	}
	gID, _ := strconv.Atoi(gameID)

	if err := g.store.JoinGame(&models.PlayersInGame{
		GameID:   uint(gID),
		PlayerID: player.ID,
	}); err != nil {
		return nil, err
	}

	return &entities.Game{
		GameID:        gameID,
		PlayersJoined: 1,
	}, nil

}
