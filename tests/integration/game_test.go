package webrung_test

import (
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

func TestGame_CanCreateNewWithAPI(t *testing.T) {

	c := http.Client{}
	contentType := "application/json"
	reqURI := fmt.Sprintf("%s/api/v1/games", API_URL)
	resp, err := c.Post(reqURI, contentType, nil)
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var gr GameResponse
	assert.Nil(t, json.Unmarshal(b, &gr))
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
