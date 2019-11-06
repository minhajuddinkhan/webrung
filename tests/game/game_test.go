package webrung_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PORT = os.Getenv("PORT")
var HOST = os.Getenv("HOST")
var API_URL = fmt.Sprintf("http://%s:%s", HOST, PORT)

type GameResponse struct {
	GameID        string `json:"game_id"`
	PlayersJoined int    `json:"players_joined"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

func GetAuthToken(username string) (string, error) {

	jsonBody := []byte(fmt.Sprintf(`{"username": "%s"}`, username))
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

	var lr LoginResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&lr)
	return lr.Token, err
}

func TestGame_ShouldBeCreated(t *testing.T) {

	token, err := GetAuthToken("North")
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
	if err != nil {
		t.Fail()
		return
	}
	assert.Nil(t, err)
	dec := json.NewDecoder(resp.Body)

	gr := GameResponse{}
	err = dec.Decode(&gr)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotZero(t, len(gr.GameID))
	assert.Equal(t, 1, gr.PlayersJoined)
}

func TestGame_ShouldGetByID(t *testing.T) {

	token, err := GetAuthToken("North")
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

	dec := json.NewDecoder(resp.Body)
	gr := GameResponse{}
	err = dec.Decode(&gr)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	c = http.Client{}
	resp, err = c.Get(fmt.Sprintf("%s/api/v1/games/%s", API_URL, gr.GameID))
	assert.Nil(t, err)

	dec = json.NewDecoder(resp.Body)
	var getGameResponse GameResponse
	err = dec.Decode(&getGameResponse)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.Nil(t, err)
	assert.Equal(t, gr.GameID, getGameResponse.GameID)
}

func TestGame_ShouldAllowJoiningIt(t *testing.T) {

	token, err := GetAuthToken("North")
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

	dec := json.NewDecoder(resp.Body)
	gr := GameResponse{}
	err = dec.Decode(&gr)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	token, err = GetAuthToken("South")
	assert.Nil(t, err)

	c = http.Client{}
	resp, err = c.Get(fmt.Sprintf("%s/api/v1/games/%s/join", API_URL, gr.GameID))
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGame_ShouldNotAllowJoiningItWithSamePlayer(t *testing.T) {

	token, err := GetAuthToken("North")
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

	dec := json.NewDecoder(resp.Body)
	gr := GameResponse{}
	err = dec.Decode(&gr)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	token, err = GetAuthToken("North")
	assert.Nil(t, err)

	c = http.Client{}
	url = fmt.Sprintf("%s/api/v1/games/%s/join", API_URL, gr.GameID)
	r, err = http.NewRequest(http.MethodGet, url, nil)
	assert.Nil(t, err)
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err = c.Do(r)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.StatusCode, 400)
}
