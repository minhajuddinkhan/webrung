package models

import (
	"github.com/jinzhu/gorm"
)

//Card card model db
type Card struct {
	gorm.Model
	House  string
	Number int
}
