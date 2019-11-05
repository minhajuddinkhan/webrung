package controllers

import (
	"encoding/json"
	"net/http"

	boom "github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	pm "github.com/minhajuddinkhan/webrung/managers/player"
	"github.com/minhajuddinkhan/webrung/store"
)

//GetPlayer GetPlayer
func GetPlayer(store store.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := mux.Vars(r)["id"]
		manager := pm.NewPlayerManager(store)
		newGame, err := manager.GetPlayer(playerID)
		if err != nil {
			switch err.(type) {
			case (*errors.ErrPlayerNotFound):
				boom.NotFound(w, err.Error())
				return
			}

			boom.Internal(w)
			return
		}

		w.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(newGame); err != nil {
			boom.Internal(w)
			return
		}

	}
}

//CreatePlayer CreatePlayer
func CreatePlayer(store store.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		mgr := pm.NewPlayerManager(store)
		dec := json.NewDecoder(r.Body)
		var p entities.Player
		if err := dec.Decode(&p); err != nil {
			boom.BadRequest(w, err)
			return
		}
		newPlayer, err := mgr.CreatePlayer(&p)
		if err != nil {
			switch err.(type) {

			case (*errors.ErrFailCreatePlayerInDb):
				w.WriteHeader(500)
				return
			}
			w.WriteHeader(500)
			return
		}

		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(newPlayer); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return

	}
}
