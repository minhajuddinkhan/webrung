package webrung

type DB struct {
	ConnectionString string
	Dialect          string
}

//Conf webrung conf
type Conf struct {
	DB DB
}
