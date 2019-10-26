package models

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Players []Player
	Winners []Player
}
