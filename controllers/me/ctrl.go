package me

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

type Controller struct {
	playerStore store.Player
	gameStore   store.Game
	ioclient    iorpc.Client
}

//NewMeController NewMeController
func NewMeController(ioclient iorpc.Client, playerStore store.Player, gameStore store.Game) Controller {

	return Controller{
		ioclient:    ioclient,
		playerStore: playerStore,
		gameStore:   gameStore,
	}
}
