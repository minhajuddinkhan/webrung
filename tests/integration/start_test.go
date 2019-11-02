package webrung_test

import (
	"fmt"
	"os"
)

var PORT = os.Getenv("PORT")
var HOST = os.Getenv("HOST")
var API_URL = fmt.Sprintf("http://%s:%s", HOST, PORT)
