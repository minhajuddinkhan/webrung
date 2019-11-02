package webrung

type DB struct {
	ConnectionString string
	Dialect          string
}

type AuthRedis struct {
	RedisURL string
}

//Conf webrung conf
type Conf struct {
	DB        DB
	AuthRedis AuthRedis
}
