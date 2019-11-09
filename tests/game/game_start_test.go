package webrung_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StartGame(gameID, token string) (int, error) {

	c := http.Client{}
	url := fmt.Sprintf("%s/api/v1/games/%s/start", API_URL, gameID)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return r.Response.StatusCode, err
	}
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	if err != nil {
		return resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("start game request failed")
	}
	return resp.StatusCode, nil

}

func TestGame_ShouldStart(t *testing.T) {

	t1, err := GetAuthToken("North")
	assert.Nil(t, err)

	t2, err := GetAuthToken("East")
	assert.Nil(t, err)

	t3, err := GetAuthToken("West")
	assert.Nil(t, err)

	t4, err := GetAuthToken("South")
	assert.Nil(t, err)

	game, err := CreateGame(t1)
	if err != nil {
		t.Errorf("game creation failed. err: %v", err)
		return
	}

	assert.Nil(t, JoinGame(t2, game.GameID))
	assert.Nil(t, JoinGame(t3, game.GameID))
	assert.Nil(t, JoinGame(t4, game.GameID))

	statusCode, err := StartGame(game.GameID, t1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

}

func TestGame_ShouldFailToStartOnLessPlayersJoined(t *testing.T) {

	t1, err := GetAuthToken("North")
	assert.Nil(t, err)

	t2, err := GetAuthToken("East")
	assert.Nil(t, err)

	t3, err := GetAuthToken("West")
	assert.Nil(t, err)

	game, err := CreateGame(t1)
	if err != nil {
		t.Errorf("game creation failed. err: %v", err)
		return
	}

	assert.Nil(t, JoinGame(t2, game.GameID))
	assert.Nil(t, JoinGame(t3, game.GameID))

	statusCode, err := StartGame(game.GameID, t1)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)

}

func TestGame_ShouldFailOnStartBySomeoneOtherThanHost(t *testing.T) {

	t1, err := GetAuthToken("North")
	assert.Nil(t, err)

	t2, err := GetAuthToken("East")
	assert.Nil(t, err)

	t3, err := GetAuthToken("West")
	assert.Nil(t, err)

	game, err := CreateGame(t1)
	if err != nil {
		t.Errorf("game creation failed. err: %v", err)
		return
	}

	assert.Nil(t, JoinGame(t2, game.GameID))
	assert.Nil(t, JoinGame(t3, game.GameID))

	statusCode, err := StartGame(game.GameID, t2)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)

}
