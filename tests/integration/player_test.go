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

	b, _ := json.Marshal(p)
	reader := bytes.NewReader(b)

	resp, err := c.Post(reqURI, contentType, reader)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	b, err = ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	var pr PlayerResponse
	assert.Nil(t, json.Unmarshal(b, &pr))
	assert.NotNil(t, resp)

	playerID = pr.ID
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
