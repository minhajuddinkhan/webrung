package game

import (
	"net/http"
	"strconv"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) JoinGame(w http.ResponseWriter, r *http.Request) {

	gameID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		boom.BadRequest(w, "invalid game id")
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
		&entities.Game{GameID: uint(gameID)})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
	_, err = ctrl.ioclient.SetGameIDInToken(iorpc.JoinGameRequest{
		GameID: uint(gameID),
		Token:  r.Header.Get("token"),
	})

	if err != nil {
		boom.BadRequest(w, err)
		return
	}
}
