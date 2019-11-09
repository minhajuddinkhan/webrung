package game

import (
	"net/http"

	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

//Controller game controller
type Controller interface {
	CreateGame(w http.ResponseWriter, r *http.Request)
	JoinGame(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
	StartGame(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	gameStore store.Game
	ioclient  iorpc.Client
}

//NewGameCtrl returns a new game controller
func NewGameCtrl(gameStore store.Game, ioclient iorpc.Client) Controller {
	return &controller{gameStore: gameStore, ioclient: ioclient}
}
