package game

import (
	"net/http"

	gm "github.com/minhajuddinkhan/webrung/managers/game"
)

func (ctrl *controller) StartGame(w http.ResponseWriter, r *http.Request) {

	gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
}
