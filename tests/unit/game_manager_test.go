package game_test

import (
	"testing"

	gm "github.com/minhajuddinkhan/webrung/managers/game"
	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/stretchr/testify/assert"
)

func beforeEach() gm.GameManager {
	store, _ := mocks.NewStore()
	gm := gm.NewGameManager(store)
	return gm
}

func TestGameManger_CanCreateGame(t *testing.T) {

	manager := beforeEach()
	game, err := manager.CreateGame()
	assert.Nil(t, err)
	assert.Equal(t, "123", game.GameID)
	assert.Equal(t, 0, game.PlayersJoined)

}

func TestGameManager_CanGetGame(t *testing.T) {

	manager := beforeEach()
	game, err := manager.GetGame("69")
	assert.Nil(t, err)
	assert.NotNil(t, game)

	assert.Equal(t, game.GameID, "69")
	assert.Equal(t, game.PlayersJoined, 1)
}
