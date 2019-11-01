package webrung_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var PORT = os.Getenv("PORT")
var HOST = os.Getenv("HOST")
var API_URL = fmt.Sprintf("http://%s:%s", HOST, PORT)

type PlayerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var playerID string

func TestPlayer_CanCreateNewWithAPI(t *testing.T) {

	c := http.Client{}
	contentType := "application/json"
	reqURI := fmt.Sprintf("%s/api/v1/players", API_URL)
	p := PlayerResponse{
		Name: "Tesla",
	}
	resp, err := c.Post(reqURI, contentType, p)
	assert.Nil(t, err)
	if err != nil {
		spew.Dump(err.Error())
	}
	assert.NotNil(t, resp)

	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var pr PlayerResponse
	assert.Nil(t, json.Unmarshal(b, &pr))
	assert.NotNil(t, resp)

	assert.Equal(t, 0, gr.PlayersJoined)
	gameID = gr.GameID
}

func TestPlayer_CanGetWithAPI(t *testing.T) {

	c := http.Client{}
	resp, err := c.Get(fmt.Sprintf("%s/api/v1/players/%s", API_URL, playerID))
	assert.Nil(t, err)
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var pr PlayerResponse
	err = json.Unmarshal(bytes, &pr)
	assert.Nil(t, err)
	assert.Equal(t, pr.ID, pr.ID)
	assert.Equal(t, pr.Name, pr.Name)

}
