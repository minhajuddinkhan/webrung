package webrung_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func JoinGame(token string, gameID string) error {
	c := http.Client{}
	url := fmt.Sprintf("%s/api/v1/games/%s/join", API_URL, gameID)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("join game request failed")
	}

	return nil
}

func TestGame_ShouldAllowJoiningIt(t *testing.T) {

	t1, err := GetAuthToken("North")
	assert.Nil(t, err)

	game, err := CreateGame(t1)
	assert.Nil(t, err)

	t2, err := GetAuthToken("South")
	assert.Nil(t, err)

	err = JoinGame(t2, game.GameID)
	assert.Nil(t, err)

}

func TestGame_ShouldNotAllowJoiningItWithSamePlayer(t *testing.T) {

	t1, err := GetAuthToken("North")
	game, err := CreateGame(t1)
	assert.Nil(t, err)

	err = JoinGame(t1, game.GameID)
	assert.NotNil(t, err)

}
