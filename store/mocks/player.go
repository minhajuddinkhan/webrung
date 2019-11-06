package mocks

import (
	"fmt"
	"strconv"

	"github.com/minhajuddinkhan/webrung/store/models"
)

//CreatePlayer CreatePlayer
func (m *PlayerStore) CreatePlayer(p *models.Player) (string, error) {
	m.player = *p

	return strconv.FormatUint(uint64(m.player.ID), 10), nil
}

//GetPlayer GetPlayer
func (m *PlayerStore) GetPlayer(playerID string) (*models.Player, error) {
	return &m.player, nil
}

//GetPlayerByName GetPlayerByName
func (m *PlayerStore) GetPlayerByName(name string) (*models.Player, error) {
	if m.connErr {
		return nil, fmt.Errorf("mock error from db")
	}
	p := &models.Player{}
	p.SetID("1")
	p.Name = "mockPlayer"
	return p, nil
}
