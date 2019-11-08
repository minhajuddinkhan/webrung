package router

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

//Mux mux router
type Mux interface {
	Router() *mux.Router
	RegisterGameRoutes(store.Game, iorpc.Client)
	RegisterPlayerRoutes(store.Player)
	RegisterAuthRoutes(store.Player, iorpc.Client)
}

type router struct {
	router *mux.Router
}

//New creates a new router
func New() Mux {
	return &router{router: mux.NewRouter()}
}

func (r *router) Router() *mux.Router {
	return r.router
}
