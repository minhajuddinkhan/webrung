package game

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) JoinGame(w http.ResponseWriter, r *http.Request) {

	gameID := mux.Vars(r)["id"]
	spew.Dump("GameID?", gameID)
	if gameID == "" {
		boom.BadRequest(w, "empty game id")
		return
	}

	token := r.Header.Get("token")
	_, playerID, err := ctrl.ioclient.Authenticate(token)
	if err != nil {
		boom.Unathorized(w, err)
		return
	}

	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)

	err = gameManager.JoinGame(
		&entities.Player{ID: playerID},
		&entities.Game{GameID: gameID})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
	_, err = ctrl.ioclient.SetGameIDInToken(iorpc.JoinGameRequest{
		GameID: gameID,
		Token:  r.Header.Get("token"),
	})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
