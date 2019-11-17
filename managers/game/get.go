package game

import (
	"github.com/minhajuddinkhan/webrung/entities"
)

func (g *gameManager) GetGame(gameID uint) (*entities.Game, error) {
	game, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err

	}

	players, err := g.store.GetPlayersInGame(gameID)
	if err != nil {
		return nil, err
	}

	return &entities.Game{
		GameID:        game.ID,
		PlayersJoined: len(players),
		HostID:        game.HostID,
	}, nil

}

func (g *gameManager) GetJoinableGames(requestedByPlayerID uint) ([]entities.Game, error) {

	joinableGames, err := g.store.GetJoinableGames()
	if err != nil {
		return nil, err
	}

	games := []entities.Game{}
	for _, game := range joinableGames {

		// cant view a game which a player has already joined.
		if game.PlayerID == requestedByPlayerID {
			continue
		}
		games = append(games, entities.Game{
			GameID:        game.GameID,
			PlayersJoined: game.PlayersJoined,
		})
	}
	return games, nil
}
