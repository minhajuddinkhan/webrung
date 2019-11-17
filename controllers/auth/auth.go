package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	boom "github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/managers/auth"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

//Authenticate Authenticate
func (ctrl *controller) Login(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	var body LoginRequest
	err := dec.Decode(&body)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	mgr := auth.NewAuthManager(ctrl.ioclient, ctrl.playerStore, ctrl.gameStore)
	token, err := mgr.Login(body.Username)
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

func (ctrl *controller) Logout(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("token")
	if token == "" {
		boom.BadRequest(w, fmt.Errorf("token is invalid"))
		return
	}

	mgr := auth.NewAuthManager(ctrl.ioclient, ctrl.playerStore, ctrl.gameStore)
	err := mgr.Logout(token)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
