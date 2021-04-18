package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var path = os.Getenv("W0Q_DB_PATH")
	if len(path) == 0 {
		path = "database.sqlite"
	}
	database, err := gorm.Open("sqlite3", path)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	database.AutoMigrate(&Url{})

	DB = database
}
