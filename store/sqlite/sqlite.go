package sqlite

import (
	//sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/webrung/store"
)

//NewPlayerStore NewPlayerStore
func NewPlayerStore(connStr string) store.Player {
	return &playerStore{connStr: connStr, dialect: "sqlite3"}
}

//NewGameStore NewGameStore
func NewGameStore(connStr string) store.Game {
	return &game{connStr: connStr, dialect: "sqlite3"}
}
