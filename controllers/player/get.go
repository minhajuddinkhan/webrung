package player

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/errors"
	pm "github.com/minhajuddinkhan/webrung/managers/player"
)

//GetPlayer GetPlayer
func (ctrl *controller) GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := mux.Vars(r)["id"]
	manager := pm.NewPlayerManager(ctrl.playerStore)
	newGame, err := manager.GetPlayer(playerID)
	if err != nil {
		switch err.(type) {
		case (*errors.ErrPlayerNotFound):
			boom.NotFound(w, err.Error())
			return
		}

		boom.Internal(w)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(newGame); err != nil {
		boom.Internal(w)
		return
	}

}
