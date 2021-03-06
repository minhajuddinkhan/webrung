package game

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/errors"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) GetAllGames(w http.ResponseWriter, r *http.Request) {

	_, playerID, err := ctrl.ioclient.Authenticate(r.Header.Get(AuthHeader))
	if err != nil {
		boom.Unathorized(w, err)
		return
	}

	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	games, err := gameManager.GetJoinableGames(playerID)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(&games); err != nil {
		boom.BadRequest(w, err)
		return
	}
}

func (ctrl *controller) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	newGame, err := gameManager.GetGame(uint(gameID))
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
