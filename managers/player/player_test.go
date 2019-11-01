package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/minhajuddinkhan/webrung/entities"
	pm "github.com/minhajuddinkhan/webrung/managers/player"
	"github.com/minhajuddinkhan/webrung/store/mocks"
)

func beforeEach(errStore bool) pm.Manager {
	store, _ := mocks.NewStore(errStore)
	pm := pm.NewPlayerManager(store)
	return pm
}

var playerID string

func TestPlayer_CanCreate(t *testing.T) {

	shouldErr := false
	mgr := beforeEach(shouldErr)

	p := entities.Player{Name: "Naruto"}
	player, err := mgr.CreatePlayer(&p)
	assert.Nil(t, err)
	assert.NotNil(t, player)
	playerID = player.ID
}

func TestPlayer_CanGet(t *testing.T) {

	shouldErr := false
	mgr := beforeEach(shouldErr)
	player, err := mgr.GetPlayer(playerID)
	assert.Nil(t, err)
	assert.Equal(t, player.ID, playerID)

}
