package me

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
	pm "github.com/minhajuddinkhan/webrung/managers/player"
)

//Info Info
func (ctrl *Controller) Info(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("token")
	gameID, playerID, err := ctrl.ioclient.Authenticate(token)
	if err != nil {
		boom.Unathorized(w, err)
		return
	}
	fmt.Println(gameID, playerID)

	playerMgr := pm.NewPlayerManager(ctrl.playerStore)
	player, err := playerMgr.GetPlayer(playerID)
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	enc := json.NewEncoder(w)
	if gameID == "" || gameID == "0" {
		if err := enc.Encode(player); err != nil {
			boom.Internal(w)
			return
		}
		return
	}

	gameManager := gm.NewGameManager(ctrl.gameStore, ctrl.ioclient)
	game, err := gameManager.GetGame(gameID)
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	player.GameID = game.GameID
	player.InGame = true
	if game.HostID == player.ID {
		player.IsHost = true
	}

	if err := enc.Encode(player); err != nil {
		boom.Internal(w, err)
		return
	}

}
