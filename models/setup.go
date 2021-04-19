package models

import (
	"github.com/Webbaum/w0q/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", config.MyConfig.DBPath)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	database.AutoMigrate(&Url{})

	DB = database
}
