package game

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) StartGame(w http.ResponseWriter, r *http.Request) {

	gameID := mux.Vars(r)["id"]
	mgr := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	err := mgr.StartGame(gameID)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
