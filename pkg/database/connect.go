package database

import (
	"github.com/glebarez/sqlite"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to sqlite database!", err)
	}

	DB = conn
}

func Connection() *gorm.DB {
	return DB
}
