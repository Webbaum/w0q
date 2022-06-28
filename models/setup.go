package models

import (
	"github.com/Webbaum/w0q/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var con gorm.Dialector
	switch config.MyConfig.DBDialect {
	case "sqlite3":
		con = sqlite.Open(config.MyConfig.DBPath)
	case "postgres":
		con = postgres.Open(config.MyConfig.DBPath)
	default:
		panic("Invalid sql dialect provided")
	}
	database, err := gorm.Open(con)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	database.AutoMigrate(&Url{})

	DB = database
}
