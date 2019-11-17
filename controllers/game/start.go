package game

import (
	"net/http"
	"strconv"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/entities"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) StartGame(w http.ResponseWriter, r *http.Request) {

	gameID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		boom.BadRequest(w, err)
		return
	}
	token := r.Header.Get(AuthHeader)

	_, playerID, err := ctrl.ioclient.Authenticate(token)
	if err != nil {
		boom.Unathorized(w)
		return
	}

	mgr := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	_, err = mgr.StartGame(uint(gameID), &entities.Player{
		ID: playerID,
	})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
