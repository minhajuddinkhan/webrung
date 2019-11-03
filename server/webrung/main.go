package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/webrung"
	"github.com/minhajuddinkhan/webrung/controllers"
	"github.com/minhajuddinkhan/webrung/store"
)

func main() {

	fmt.Println("here bro..")

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

	ioRungClient, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%s", conf.IORung.Host, conf.IORung.Port))
	if err != nil {
		log.Fatal(err)
	}

	store, err := store.NewRungStore(conf.DB.Dialect, conf.DB.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Migrate(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	// Game REST
	r.HandleFunc("/api/v1/games", controllers.CreateGame(store)).Methods("POST")
	r.HandleFunc("/api/v1/games/{id}", controllers.GetGame(store)).Methods("GET")

	// Player REST
	r.HandleFunc("/api/v1/players", controllers.CreatePlayer(store)).Methods("POST")
	r.HandleFunc("/api/v1/players/{id}", controllers.GetPlayer(store)).Methods("GET")

	r.HandleFunc("/api/v1/auth", controllers.Authenticate(ioRungClient, store)).Methods("POST")
	http.Handle("/", r)

	spew.Dump("LISTENING ON PORT", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
}
