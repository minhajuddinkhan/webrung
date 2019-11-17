package auth

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (m *manager) Login(username string) (token string, err error) {
	player, err := m.playerStore.GetPlayerByName(username)
	if err != nil {
		return "", err
	}

	game, err := m.gameStore.GetGameByPlayer(player.ID)
	if err != nil {
		return "", err
	}

	req := iorpc.AddPlayerRequest{
		PlayerID: player.ID,
		GameID:   game.ID,
	}
	return m.ioclient.AddPlayer(req)

}

func (m *manager) Logout(token string) error {

	_, err := m.ioclient.Logout(token)
	return err
}
