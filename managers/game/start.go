package game

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (gm *gameManager) StartGame(gameID string) error {

	// gameToStart, err := gm.store.GetGame(gameID)
	// if err != nil {
	// 	return err
	// }

	//TODO:: finalize architecture and implement start game
	// if gameToStart.PlayersJoined != 4 {
	// 	return fmt.Errorf("unable to start game: all players not joined yet")
	// }

	players, err := gm.store.GetPlayersInGame(gameID)
	if err != nil {
		return err
	}
	playerIds := make([]string, len(players))
	for i, p := range players {
		playerIds[i] = p.GetPlayerID()
	}

	_, err = gm.ioclient.StartGame(iorpc.DistributeCardsRequest{
		PlayerIds: playerIds,
		GameID:    gameID,
	})
	return err
}
