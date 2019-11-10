package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/minhajuddinkhan/webrung/config"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/router"
	"github.com/minhajuddinkhan/webrung/store"
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

	fmt.Println("LISTENING ON PORT", conf.HTTPPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTPPort), r.Handler()))
}
