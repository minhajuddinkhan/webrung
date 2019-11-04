package webrung_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GameResponse struct {
	GameID        string `json:"game_id"`
	PlayersJoined int    `json:"players_joined"`
}

var gameID string

func GetAuthToken() (string, error) {

	jsonBody := []byte(`{"username": "North"}`)
	url := fmt.Sprintf("%s/api/v1/auth", API_URL)
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	var lr loginResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&lr)
	return lr.Token, err
}

func TestGame_ShouldBeCreated(t *testing.T) {

	token, err := GetAuthToken()
	assert.Nil(t, err)

	c := http.Client{}
	url := fmt.Sprintf("%s/api/v1/games", API_URL)
	r, err := http.NewRequest(http.MethodPost, url, nil)
	assert.Nil(t, err)
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	assert.Nil(t, err)
	dec := json.NewDecoder(resp.Body)

	gr := GameResponse{}
	err = dec.Decode(&gr)

	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.NotZero(t, len(gr.GameID))
	assert.Equal(t, 1, gr.PlayersJoined)
	gameID = gr.GameID
}

func TestGame_CanGetWithAPI(t *testing.T) {

	c := http.Client{}
	resp, err := c.Get(fmt.Sprintf("%s/api/v1/games/%s", API_URL, gameID))
	assert.Nil(t, err)
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var gr GameResponse
	err = json.Unmarshal(bytes, &gr)
	assert.Nil(t, err)
	assert.Equal(t, gr.GameID, gameID)
}

func TestGame_CanJoin(t *testing.T) {
	c := http.Client{}
	resp, err := c.Get(fmt.Sprintf("%s/api/v1/games/%s", API_URL, gameID))
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
