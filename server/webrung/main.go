package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung"
	"github.com/minhajuddinkhan/webrung/controllers"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

func main() {

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("empty http port from PORT env")
	}

	dbConnectionString := os.Getenv("DB_URL")
	if dbConnectionString == "" {
		log.Fatal("empty db connection string from DB_URL")
	}

	dialect := os.Getenv("DB_DIALECT")
	if dialect == "" {
		log.Fatal("empty db dialect from DB_DIALECT")
	}

	ioRungHost := os.Getenv("IORUNG_HOST")
	if ioRungHost == "" {
		log.Fatal("empty iorung host")
	}

	ioRungPort := os.Getenv("IORUNG_PORT")
	if ioRungHost == "" {
		log.Fatal("empty iorung port")
	}

	conf := webrung.Conf{
		DB: webrung.DB{
			ConnectionString: dbConnectionString,
			Dialect:          dialect,
		},
		IORung: webrung.IORung{
			Host: ioRungHost,
			Port: ioRungPort,
		},
	}

	client, err := iorpc.NewIOClient(&conf)
	if err != nil {
		log.Fatal(err)
	}

	gameStore, err := store.NewGameStore(conf.DB.Dialect, conf.DB.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	playerStore, err := store.NewPlayerStore(conf.DB.Dialect, conf.DB.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	// // Game REST
	r.HandleFunc("/api/v1/games", controllers.CreateGame(gameStore, client)).Methods("POST")
	r.HandleFunc("/api/v1/games/{id}", controllers.GetGame(gameStore, client)).Methods("GET")
	r.HandleFunc("/api/v1/games/{id}/join", controllers.JoinGame(gameStore, client)).Methods("GET")

	// Player REST
	r.HandleFunc("/api/v1/players", controllers.CreatePlayer(playerStore)).Methods("POST")
	r.HandleFunc("/api/v1/players/{id}", controllers.GetPlayer(playerStore)).Methods("GET")

	r.HandleFunc("/api/v1/auth", controllers.Authenticate(client, playerStore)).Methods("POST")
	http.Handle("/", r)

	spew.Dump("LISTENING ON PORT", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
}
