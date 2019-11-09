package game

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/entities"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) StartGame(w http.ResponseWriter, r *http.Request) {

	gameID := mux.Vars(r)["id"]
	token := r.Header.Get(AuthHeader)

	_, playerID, err := ctrl.ioclient.Authenticate(token)
	if err != nil {
		boom.Unathorized(w)
		return
	}

	mgr := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	_, err = mgr.StartGame(gameID, &entities.Player{
		ID: playerID,
	})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
