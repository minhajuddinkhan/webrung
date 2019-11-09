package game

import (
	"github.com/minhajuddinkhan/webrung/entities"
)

func (g *gameManager) GetGame(gameID string) (*entities.Game, error) {
	game, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err

	}

	players, err := g.store.GetPlayersInGame(gameID)
	if err != nil {
		return nil, err
	}

	return &entities.Game{
		GameID:        game.GetID(),
		PlayersJoined: len(players),
	}, nil

}
