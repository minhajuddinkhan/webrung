package webrung_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type reqBody struct {
	username string
}

type loginResponse struct {
	Token string `json:"token,omitempty"`
}

var PORT = os.Getenv("PORT")
var HOST = os.Getenv("HOST")
var API_URL = fmt.Sprintf("http://%s:%s", HOST, PORT)

func TestAuth_ShouldPlayerLogin(t *testing.T) {
	c := http.Client{}
	contentType := "application/json"
	reqURI := fmt.Sprintf("%s/api/v1/auth", API_URL)

	jsonBody := []byte(`{"username": "North"}`)
	resp, err := c.Post(reqURI, contentType, bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	var lr loginResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&lr)
	assert.Nil(t, err)
	//is a jwt token
	assert.Equal(t, 3, len(strings.Split(lr.Token, ".")))
}

func TestAuth_ShouldFailLoginOnBadCreds(t *testing.T) {

	c := http.Client{}
	contentType := "application/json"
	reqURI := fmt.Sprintf("%s/api/v1/auth", API_URL)

	jsonBody := []byte(`{"username": "GameOfThrones"}`)
	resp, err := c.Post(reqURI, contentType, bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusUnauthorized)
}
