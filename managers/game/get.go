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

func (g *gameManager) GetJoinableGames() ([]entities.Game, error) {

	joinableGames, err := g.store.GetJoinableGames()
	if err != nil {
		return nil, err
	}

	games := make([]entities.Game, len(joinableGames))
	for j, game := range joinableGames {
		games[j].GameID = game.GameID
		games[j].PlayersJoined = game.PlayersJoined
	}
	return games, nil
}
