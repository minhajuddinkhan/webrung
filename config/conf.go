package config

import (
	"log"
	"os"
)

type DB struct {
	ConnectionString string
	Dialect          string
}

type Redis struct {
	RedisURL string
}

type IORung struct {
	Host string
	Port string
}

//Conf webrung conf
type Conf struct {
	DB          DB
	AuthRedis   Redis
	SocketRedis Redis
	IORung      IORung
	HTTPPort    string
}

func New() Conf {

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

	return Conf{
		HTTPPort: httpPort,
		DB: DB{
			ConnectionString: dbConnectionString,
			Dialect:          dialect,
		},
		IORung: IORung{
			Host: ioRungHost,
			Port: ioRungPort,
		},
	}
}
