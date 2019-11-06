package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/errors"
	"github.com/minhajuddinkhan/webrung/iorpc"
	gm "github.com/minhajuddinkhan/webrung/managers/game"
	"github.com/minhajuddinkhan/webrung/store"
)

//AuthHeader AuthHeader
var AuthHeader = "token"

//GetGame GetGame
func GetGame(store store.Game, c iorpc.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID := mux.Vars(r)["id"]
		gameManager := gm.NewGameManager(store, c)
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
func CreateGame(store store.Game, c iorpc.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		gameManager := gm.NewGameManager(store, c)
		token := r.Header.Get(AuthHeader)
		_, playerID, err := c.Authenticate(token)
		if err != nil {
			boom.Unathorized(w, err)
			return
		}
		newGame, err := gameManager.CreateGame(&entities.Player{
			ID: playerID,
		})
		if err != nil {
			switch err.(type) {

			case (*errors.ErrFailCreateGameInDb):
				boom.Internal(w)
				return
			}
			boom.Internal(w)
			return
		}

		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(newGame); err != nil {
			boom.Internal(w)
			return
		}
		return

	}
}

//JoinGame JoinGame
func JoinGame(store store.Game, iorungrpc iorpc.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		gameID := mux.Vars(r)["id"]
		if gameID == "" {
			boom.BadRequest(w, "empty game id")
			return
		}

		token := r.Header.Get("token")
		_, playerID, err := iorungrpc.Authenticate(token)
		if err != nil {
			boom.Unathorized(w, err)
			return
		}

		gameManager := gm.NewGameManager(store, iorungrpc)

		err = gameManager.JoinGame(
			&entities.Player{ID: playerID},
			&entities.Game{GameID: gameID})

		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		_, err = iorungrpc.SetGameIDInToken(iorpc.JoinGameRequest{
			GameID: gameID,
			Token:  r.Header.Get("token"),
		})

		if err != nil {
			boom.BadRequest(w, err)
			return
		}
	}
}
