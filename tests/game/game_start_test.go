package webrung_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/stretchr/testify/assert"
)

func CreateGame(token string) (GameResponse, error) {

	c := http.Client{}
	url := fmt.Sprintf("%s/api/v1/games", API_URL)
	r, err := http.NewRequest(http.MethodPost, url, nil)
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	var gr GameResponse
	resp, err := c.Do(r)

	if err != nil {
		return gr, err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&gr)

	if gr.GameID == "" {
		return gr, fmt.Errorf("game id is empty")
	}
	return gr, err
}

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

func StartGame(gameID, token string) ([]entities.Player, error) {

	c := http.Client{}
	url := fmt.Sprintf("%s/api/v1/games/%s/start", API_URL, gameID)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("start game request failed")
	}

	dec := json.NewDecoder(resp.Body)
	var players []entities.Player

	return players, dec.Decode(&players)

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

	players, err := StartGame(game.GameID, t1)
	assert.Nil(t, err)

	assert.Equal(t, 4, len(players))
	for _, p := range players {
		assert.Equal(t, 13, len(p.Cards))
		for _, c := range p.Cards {
			assert.NotEmpty(t, c)
		}

	}
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

	_, err = StartGame(game.GameID, t1)
	assert.NotNil(t, err)

}

//TODO:: add negative test case for start game
