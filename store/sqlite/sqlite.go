package sqlite

import (
	//sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Store Store
type Store struct {
	connStr string
	dialect string
}

//NewStore NewStore
func NewStore(connStr string) (*Store, error) {
	return &Store{connStr: connStr, dialect: "sqlite3"}, nil
}
