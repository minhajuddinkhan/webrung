package webrung_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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
	r, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	r.Header.Set("token", token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(r)
	b, _ := ioutil.ReadAll(resp.Body)
	spew.Dump(b)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("join game request failed %v", resp.StatusCode)
	}

	if err != nil {
		return err
	}

	return nil
}

// func TestGame_ShouldStart(t *testing.T) {

// 	t1, err := GetAuthToken("North")
// 	assert.Nil(t, err)

// 	t2, err := GetAuthToken("East")
// 	assert.Nil(t, err)

// 	t3, err := GetAuthToken("West")
// 	assert.Nil(t, err)

// 	t4, err := GetAuthToken("South")
// 	assert.Nil(t, err)

// 	game, err := CreateGame(t1)
// 	if err != nil {
// 		t.Errorf("game creation failed. err: %v", err)
// 		return
// 	}

// 	assert.Nil(t, JoinGame(t2, game.GameID))
// 	assert.Nil(t, JoinGame(t3, game.GameID))
// 	assert.Nil(t, JoinGame(t4, game.GameID))

// }
