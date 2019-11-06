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
	store, _ := mocks.NewPlayerStore(storeShouldError)

	manager := auth.NewAuthManager(mockIOClient, store)
	token, err := manager.Authenticate("North")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 3, len(strings.Split(token, ".")))
}

func TestAuth_ShouldErrorOnBadStoreState(t *testing.T) {

	mockIOClient := mock.NewMockIORPCClient()

	storeShouldError := true
	store, _ := mocks.NewPlayerStore(storeShouldError)
	manager := auth.NewAuthManager(mockIOClient, store)
	token, err := manager.Authenticate("North")
	assert.NotNil(t, err)
	assert.Empty(t, token)

}
