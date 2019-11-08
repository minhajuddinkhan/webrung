package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	config "github.com/minhajuddinkhan/webrung/config"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/router"
	"github.com/minhajuddinkhan/webrung/store"
	"github.com/rs/cors"
)

func main() {

	conf := config.New()
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

	r := router.New()
	r.RegisterGameRoutes(gameStore, client)
	r.RegisterPlayerRoutes(playerStore)
	r.RegisterAuthRoutes(playerStore, client)

	handler := cors.Default().Handler(r.Router())
	spew.Dump("LISTENING ON PORT", conf.HTTPPort)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTPPort), handler)
}
