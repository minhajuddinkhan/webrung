package game

import "github.com/minhajuddinkhan/webrung/entities"

func (g *gameManager) GetGame(gameID string) (*entities.Game, error) {
	game, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err

	}
	gameEntity := entities.Game{}
	gameEntity.Marshal(game)
	return &gameEntity, nil

}
