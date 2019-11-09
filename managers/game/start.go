package game

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (gm *gameManager) StartGame(gameID string, startBy *entities.Player) ([]entities.Player, error) {

	gameToStart, err := gm.store.GetGame(gameID)
	if err != nil {
		return nil, err
	}

	if gameToStart.GetHostID() != startBy.ID {
		return nil, fmt.Errorf("game cannot be started by someone other than the host")
	}

	players, err := gm.store.GetPlayersInGame(gameID)
	if err != nil {
		return nil, err
	}

	if len(players) != 4 {
		return nil, fmt.Errorf("cannot start game until 4 players have joined")
	}

	playerIds := make([]string, len(players))
	for i, p := range players {
		playerIds[i] = p.GetPlayerID()
	}

	resp, err := gm.ioclient.StartGame(iorpc.DistributeCardsRequest{
		PlayerIds: playerIds,
		GameID:    gameID,
	})

	var respPlayers []entities.Player

	for _, p := range resp.Players {
		var cards []entities.Card
		for _, c := range p.Cards {
			cards = append(cards, entities.Card{House: c.House, Number: c.Number})
		}
		respPlayers = append(respPlayers, entities.Player{
			Cards: cards,
			ID:    p.PlayerID,
		})
	}

	return respPlayers, err
}
