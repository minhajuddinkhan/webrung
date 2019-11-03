package webrung

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
}
