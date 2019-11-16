package auth

import (
	"net/http"

	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

//Controller auth controller
type Controller interface {
	Authenticate(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	playerStore store.Player
	gameStore   store.Game
	ioclient    iorpc.Client
}

//NewAuthController NewAuthController
func NewAuthController(store store.Player, gameStore store.Game, iorungrpc iorpc.Client) Controller {
	return &controller{playerStore: store, ioclient: iorungrpc, gameStore: gameStore}
}
