package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"

	"github.com/gorilla/mux"

	"github.com/minhajuddinkhan/webrung/errors"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
	"github.com/minhajuddinkhan/webrung/store"
)

//GetGame GetGame
func GetGame(store store.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID := mux.Vars(r)["id"]
		gameManager := gm.NewGameManager(store)
		newGame, err := gameManager.GetGame(gameID)
		if err != nil {
			switch err.(type) {
			case (*errors.ErrGameIDNotFound):
				boom.NotFound(w, err.Error())
				return
			}

			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(newGame); err != nil {
			boom.Internal(w)
			return
		}

	}
}

//CreateGame CreateGame
func CreateGame(store store.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		gameManager := gm.NewGameManager(store)
		newGame, err := gameManager.CreateGame()
		if err != nil {
			switch err.(type) {

			case (*errors.ErrFailCreateGameInDb):
				w.WriteHeader(500)
				return
			}
			w.WriteHeader(500)
			return
		}

		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(newGame); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return

	}
}
