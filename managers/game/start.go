package game

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/entities"
)

func (gm *gameManager) StartGame(player *entities.Player, gameID string) error {

	gameToStart, err := gm.store.GetGame(gameID)
	if err != nil {
		return err
	}

	//TODO:: finalize architecture and implement start game
	if gameToStart.PlayersJoined != 4 {
		return fmt.Errorf("unable to start game: all players not joined yet")
	}

	return nil
}
