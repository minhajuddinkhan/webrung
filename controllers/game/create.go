package game

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
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
		}
		boom.Internal(w)
		return
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	if err = encoder.Encode(newGame); err != nil {
		boom.Internal(w)
		return
	}
	return

}
