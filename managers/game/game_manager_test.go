package game_test

import (
	"testing"

	"github.com/minhajuddinkhan/webrung/entities"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/stretchr/testify/assert"
)

func beforeEach(errStore bool) gm.GameManager {
	store, _ := mocks.NewStore(errStore)
	gm := gm.NewGameManager(store)
	return gm
}

func TestGameManger_CanCreateGame(t *testing.T) {

	storeShouldError := false
	manager := beforeEach(storeShouldError)
	game, err := manager.CreateGame()
	assert.Nil(t, err)
	assert.Equal(t, "123", game.GameID)
	assert.Equal(t, 0, game.PlayersJoined)

}
func TestGameManager_ShouldThrowError(t *testing.T) {
	storeShouldError := true
	manager := beforeEach(storeShouldError)
	game, err := manager.CreateGame()
	assert.NotNil(t, err)
	assert.Nil(t, game)
}

func TestGameManager_CanGetGame(t *testing.T) {

	storeShouldError := false
	manager := beforeEach(storeShouldError)
	game, err := manager.GetGame("69")
	assert.Nil(t, err)
	assert.NotNil(t, game)

	assert.Equal(t, game.GameID, "69")
	assert.Equal(t, game.PlayersJoined, 1)
}

func TestGameManager_ShouldErrorOnCanGetGame(t *testing.T) {

	storeShouldError := true
	manager := beforeEach(storeShouldError)
	game, err := manager.GetGame("69")
	assert.Nil(t, game)
	assert.NotNil(t, err)

}

func TestGame_ShouldBeAbleToJoin(t *testing.T) {

	shouldRaiseError := false
	mgr := beforeEach(shouldRaiseError)
	var player entities.Player
	err := mgr.JoinGame(&player, "123")
	assert.NotNil(t, err)

	game, err := mgr.GetGame("123")
	assert.Nil(t, err)
	assert.NotNil(t, game)
}
