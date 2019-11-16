package router

import (
	"github.com/minhajuddinkhan/webrung/controllers/auth"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

func (r *router) RegisterAuthRoutes(playerStore store.Player, gameStore store.Game, ioclient iorpc.Client) {
	ctrl := auth.NewAuthController(playerStore, gameStore, ioclient)

	//Auth
	r.router.HandleFunc("/api/v1/auth", ctrl.Authenticate).Methods("POST")
}
