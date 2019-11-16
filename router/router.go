package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
	"github.com/rs/cors"
)

//Mux mux router
type Mux interface {
	Router() *mux.Router
	RegisterGameRoutes(store.Game, iorpc.Client)
	RegisterPlayerRoutes(store.Player)
	RegisterAuthRoutes(store.Player, store.Game, iorpc.Client)
	RegisterMeRoutes(playerStore store.Player, gameStore store.Game, ioclient iorpc.Client)
	Handler() http.Handler
}

type router struct {
	router  *mux.Router
	handler http.Handler
}

//New creates a new router
func New() Mux {

	r := mux.NewRouter()

	handler := cors.AllowAll().Handler(r)
	return &router{router: r, handler: handler}
}

func (r *router) Router() *mux.Router {
	return r.router
}

func (r *router) Handler() http.Handler {
	return r.handler
}
