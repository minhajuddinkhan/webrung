package player

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	pm "github.com/minhajuddinkhan/webrung/managers/player"
)

func (ctrl *controller) CreatePlayer(w http.ResponseWriter, r *http.Request) {

	mgr := pm.NewPlayerManager(ctrl.playerStore)
	dec := json.NewDecoder(r.Body)
	var p entities.Player
	if err := dec.Decode(&p); err != nil {
		boom.BadRequest(w, err)
		return
	}
	newPlayer, err := mgr.CreatePlayer(&p)
	if err != nil {
		switch err.(type) {

		case (*errors.ErrFailCreatePlayerInDb):
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(500)
		return
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	if err = encoder.Encode(newPlayer); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return

}
