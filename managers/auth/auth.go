package auth

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (m *manager) Authenticate(username string) (token string, err error) {
	player, err := m.store.GetPlayerByName(username)
	if err != nil {
		return "", err
	}
	req := iorpc.AddPlayerRequest{
		PlayerID: player.GetID(),
		GameID:   "",
	}
	return m.ioclient.AddPlayer(req)

}
