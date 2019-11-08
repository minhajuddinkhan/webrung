package game

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/errors"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID := mux.Vars(r)["id"]
	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	newGame, err := gameManager.GetGame(gameID)
	if err != nil {
		switch err.(type) {
		case (*errors.ErrGameIDNotFound):
			boom.NotFound(w, err.Error())
			return
		}

		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(newGame); err != nil {
		boom.Internal(w)
		return
	}

}
