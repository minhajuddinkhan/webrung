package router

import (
	"github.com/minhajuddinkhan/webrung/controllers/game"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

func (r *router) RegisterGameRoutes(gameStore store.Game, client iorpc.Client) {

	gameCtrl := game.NewGameCtrl(gameStore, client)

	//Game REST
	r.router.HandleFunc("/api/v1/games", gameCtrl.CreateGame).Methods("POST")
	r.router.HandleFunc("/api/v1/games/{id}", gameCtrl.GetGame).Methods("GET")
	r.router.HandleFunc("/api/v1/games/{id}/join", gameCtrl.JoinGame).Methods("GET")

}
