package models

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Winner   Player
	WinnerID uint
	Host     Player
	HostID   uint
}
