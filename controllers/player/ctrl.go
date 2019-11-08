package player

import (
	"net/http"

	"github.com/minhajuddinkhan/webrung/store"
)

//Controller http player controller
type Controller interface {

	//GetPlayer gets a player
	GetPlayer(w http.ResponseWriter, r *http.Request)

	//CreatePlayer creates a new player
	CreatePlayer(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	playerStore store.Player
}

//NewPlayerCtrl returns a new player controller
func NewPlayerCtrl(gameStore store.Player) Controller {
	return &controller{playerStore: gameStore}
}
