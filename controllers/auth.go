package controllers

import (
	"encoding/json"
	"net/http"

	boom "github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/managers/auth"
	"github.com/minhajuddinkhan/webrung/store"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

//Authenticate Authenticate
func Authenticate(iorungrpc iorpc.Client, store store.Player) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		dec := json.NewDecoder(r.Body)
		var body LoginRequest
		err := dec.Decode(&body)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}

		mgr := auth.NewAuthManager(iorungrpc, store)
		token, err := mgr.Authenticate(body.Username)
		if err != nil {
			boom.Unathorized(w, err)
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
