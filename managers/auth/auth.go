package auth

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (m *manager) Login(username string) (token string, err error) {
	player, err := m.playerStore.GetPlayerByName(username)
	if err != nil {
		return "", err
	}

	game, err := m.gameStore.GetGameByPlayer(player.GetID())
	if err != nil {
		return "", err
	}

	req := iorpc.AddPlayerRequest{
		PlayerID: player.GetID(),
		GameID:   game.GetID(),
	}
	return m.ioclient.AddPlayer(req)

}
