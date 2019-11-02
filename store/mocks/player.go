package mocks

import (
	"strconv"

	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreatePlayer CreatePlayer
func (m *Store) CreatePlayer(p *models.Player) (string, error) {
	m.player = *p

	return strconv.FormatUint(uint64(m.player.ID), 10), nil
}

//GetPlayer GetPlayer
func (m *Store) GetPlayer(playerID string) (*models.Player, error) {
	return &m.player, nil
}

func (m *Store) GetPlayerByName(name string) (*models.Player, error) {
	return nil, nil
}
