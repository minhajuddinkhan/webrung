package controllers

import (
	"encoding/json"
	"net/http"
	"net/rpc"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/store"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
}

type AddPlayerRequest struct {
	PlayerID string
	GameID   string
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

//Authenticate Authenticate
func Authenticate(iorung *rpc.Client, store store.Store) func(w http.ResponseWriter, r *http.Request) {
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
			boom.NotFound(w, "player not found")
			return
		}
		req := AddPlayerRequest{
			PlayerID: player.Name,
			GameID:   "",
		}

		var token string
		err = iorung.Call("InterfaceRPC.AddPlayer", req, &token)
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
