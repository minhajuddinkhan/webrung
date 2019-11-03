package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

//Authenticate Authenticate
func Authenticate(iorungrpc iorpc.Client, store store.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		dec := json.NewDecoder(r.Body)
		var body LoginRequest
		err := dec.Decode(&body)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		player, err := store.GetPlayerByName(body.Username)
		if err != nil {
			fmt.Println("player not found????")
			boom.NotFound(w, "player not found")
			return
		}
		fmt.Println("player??", player)
		req := iorpc.AddPlayerRequest{
			PlayerID: player.GetID(),
			GameID:   "",
		}
		token, err := iorungrpc.AddPlayer(req)
		if err != nil {
			boom.Internal(w, "")
			return
		}

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		err = enc.Encode(LoginResponse{Token: token})
		if err != nil {
			boom.Internal(w)
			return
		}
	}
}
