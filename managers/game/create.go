package game

import (
	"strconv"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store/models"
)

func (g *gameManager) CreateGame(pl *entities.Player) (*entities.Game, error) {

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

	req := iorpc.AddPlayerRequest{PlayerID: pl.ID, GameID: gameID}
	if _, err := g.ioclient.AddPlayer(req); err != nil {
		return nil, err
	}
	return &entities.Game{
		GameID:        gameID,
		PlayersJoined: 1,
	}, nil

}
