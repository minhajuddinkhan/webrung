package auth_test

import (
	"strings"
	"testing"

	"github.com/minhajuddinkhan/webrung/iorpc/mock"
	"github.com/minhajuddinkhan/webrung/managers/auth"
	"github.com/minhajuddinkhan/webrung/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuth_ShouldAuthenticate(t *testing.T) {

	mockIOClient := mock.NewMockIORPCClient()

	storeShouldError := false
	playerStore, _ := mocks.NewPlayerStore(storeShouldError)
	gameStore, _ := mocks.NewGameStore(storeShouldError)
	manager := auth.NewAuthManager(mockIOClient, playerStore, gameStore)
	token, err := manager.Login("North")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 3, len(strings.Split(token, ".")))
}

func TestAuth_ShouldErrorOnBadStoreState(t *testing.T) {

	mockIOClient := mock.NewMockIORPCClient()

	storeShouldError := true
	playerStore, _ := mocks.NewPlayerStore(storeShouldError)
	gameStore, _ := mocks.NewGameStore(storeShouldError)
	manager := auth.NewAuthManager(mockIOClient, playerStore, gameStore)
	token, err := manager.Login("North")
	assert.NotNil(t, err)
	assert.Empty(t, token)

}
