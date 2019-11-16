package router

import (
	"github.com/minhajuddinkhan/webrung/controllers/me"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

func (r *router) RegisterMeRoutes(playerStore store.Player, gameStore store.Game, ioclient iorpc.Client) {

	ctrl := me.NewMeController(ioclient, playerStore, gameStore)

	//Me
	r.router.HandleFunc("/api/v1/me", ctrl.Info).Methods("GET")
}
