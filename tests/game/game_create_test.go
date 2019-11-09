package webrung_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

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

func TestGame_ShouldBeCreated(t *testing.T) {

	token, err := GetAuthToken("North")
	assert.Nil(t, err)

	resp, err := CreateGame(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, resp.GameID)
	assert.NotEmpty(t, resp.PlayersJoined)

}
