package router

import (
	"github.com/minhajuddinkhan/webrung/controllers/player"
	"github.com/minhajuddinkhan/webrung/store"
)

func (r *router) RegisterPlayerRoutes(playerStore store.Player) {

	playerCtrl := player.NewPlayerCtrl(playerStore)

	//Player
	r.router.HandleFunc("/api/v1/players", playerCtrl.CreatePlayer).Methods("POST")
	r.router.HandleFunc("/api/v1/players/{id}", playerCtrl.GetPlayer).Methods("GET")

}
