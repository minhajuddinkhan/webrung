package game

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/iorpc"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

//AuthHeader AuthHeader
var AuthHeader = "token"

func (ctrl *controller) CreateGame(w http.ResponseWriter, r *http.Request) {

	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	token := r.Header.Get(AuthHeader)
	_, playerID, err := ctrl.ioclient.Authenticate(token)
	if err != nil {
		boom.Unathorized(w, err)
		return
	}

	newGame, err := gameManager.CreateGame(&entities.Player{
		ID: playerID,
	})
	if err != nil {
		switch err.(type) {

		case (*errors.ErrFailCreateGameInDb):
			boom.Internal(w)
			return
		case (*errors.ErrGameAlreadyHosted):
			boom.BadRequest(w, err)
			return
		case (*errors.ErrPlayerAlreadyJoinedInAnotherGame):
			boom.BadRequest(w, err)
			return
		default:
			boom.Internal(w)
			return
		}
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")

	done, err := ctrl.ioclient.SetGameIDInToken(iorpc.JoinGameRequest{
		GameID: newGame.GameID,
		Token:  token,
	})
	if !done || err != nil {
		if err := encoder.Encode(err); err != nil {
			boom.BadRequest(w, err)
		}
		return
	}

	if err = encoder.Encode(newGame); err != nil {
		boom.Internal(w)
		return
	}

}
