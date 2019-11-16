package auth

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (m *manager) Authenticate(username string) (token string, err error) {
	player, err := m.playerStore.GetPlayerByName(username)
	if err != nil {
		return "", err
	}

	fmt.Println("GETID", player.GetID())
	game, err := m.gameStore.GetGameByPlayer(player.GetID())
	if err != nil {
		return "", err
	}

	fmt.Println("GAME", game.GetID())
	req := iorpc.AddPlayerRequest{
		PlayerID: player.GetID(),
		GameID:   game.GetID(),
	}
	return m.ioclient.AddPlayer(req)

}
