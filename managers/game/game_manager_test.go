package game_test

import (
	"strconv"
	"testing"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc/mock"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/stretchr/testify/assert"
)

func beforeEach(errStore bool) gm.GameManager {
	store, _ := mocks.NewGameStore(errStore)
	client := mock.NewMockIORPCClient()
	gm := gm.NewGameManager(store, client)
	return gm
}

func TestGameManger_CanCreateGame(t *testing.T) {

	storeShouldError := false
	manager := beforeEach(storeShouldError)
	pl := entities.Player{
		ID:   "1",
		Name: "North",
	}
	game, err := manager.CreateGame(&pl)
	assert.Nil(t, err)
	assert.Equal(t, "1", game.GameID)
	assert.Equal(t, 1, game.PlayersJoined)

}
func TestGameManager_ShouldThrowError(t *testing.T) {
	storeShouldError := true
	manager := beforeEach(storeShouldError)
	pl := entities.Player{
		ID:   "1",
		Name: "North",
	}
	game, err := manager.CreateGame(&pl)
	assert.NotNil(t, err)
	assert.Nil(t, game)
}

func TestGameManager_CanGetGame(t *testing.T) {

	storeShouldError := false
	manager := beforeEach(storeShouldError)

	pl := entities.Player{
		ID:   "1",
		Name: "North",
	}
	game, err := manager.CreateGame(&pl)
	assert.Nil(t, err)
	assert.NotNil(t, game)

	outGame, err := manager.GetGame(game.GameID)
	assert.Nil(t, err)
	assert.NotNil(t, game)

	assert.Equal(t, game.GameID, outGame.GameID)
	assert.Equal(t, game.PlayersJoined, outGame.PlayersJoined)
}

func TestGameManager_ShouldErrorOnCanGetGame(t *testing.T) {

	storeShouldError := true
	manager := beforeEach(storeShouldError)
	game, err := manager.GetGame("69")
	assert.Nil(t, game)
	assert.NotNil(t, err)

}

func TestGame_ShouldBeAbleToJoin(t *testing.T) {

	storeShouldError := false
	mgr := beforeEach(storeShouldError)
	pl := entities.Player{
		ID:   "1",
		Name: "North",
	}
	game, err := mgr.CreateGame(&pl)
	assert.Nil(t, err)

	player := entities.Player{
		ID:   "1",
		Name: "North",
	}

	err = mgr.JoinGame(&player, game)
	assert.Nil(t, err)

	outGame, err := mgr.GetGame(game.GameID)
	assert.Nil(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, 2, outGame.PlayersJoined)
}

func TestGame_ShouldErrOnInvalidGameId(t *testing.T) {
	player := entities.Player{
		ID:   "1",
		Name: "North",
	}
	storeShouldError := false
	mgr := beforeEach(storeShouldError)
	game := &entities.Game{
		GameID: "abc",
	}
	err := mgr.JoinGame(&player, game)
	assert.NotNil(t, err)

}

func TestGame_ShouldErrOnInvalidPlayerId(t *testing.T) {
	player := entities.Player{
		ID:   "1abc",
		Name: "North",
	}
	storeShouldError := false
	mgr := beforeEach(storeShouldError)
	game := &entities.Game{
		GameID: "1",
	}
	err := mgr.JoinGame(&player, game)
	assert.NotNil(t, err)

}

func TestGame_ShouldThrowErrOnFiveJoins(t *testing.T) {

	storeShouldError := false
	mgr := beforeEach(storeShouldError)
	pl := entities.Player{
		ID:   "1",
		Name: "North",
	}
	game, err := mgr.CreateGame(&pl)
	assert.Nil(t, err)

	for i := 0; i < 3; i++ {
		player := entities.Player{
			ID:   strconv.Itoa(i),
			Name: "North",
		}
		err = mgr.JoinGame(&player, game)
		assert.Nil(t, err)
	}

	err = mgr.JoinGame(&entities.Player{ID: "20", Name: "John"}, game)
	assert.NotNil(t, err)

}

func TestGame_ShouldStart(t *testing.T) {

	// shouldErr := false
	// mgr := beforeEach(shouldErr)
	// mgr.StartGame()
}
