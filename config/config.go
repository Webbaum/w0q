package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
)

var MyConfig Config

func ReadConfig() {
	err := envconfig.Process("W0Q", &MyConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Config struct {
	DBDialect    string `envconfig:"W0Q_DB_DIALECT" default:"sqlite3"`
	DBPath       string `envconfig:"W0Q_DB_PATH" default:"database.sqlite"`
	LegalAddress string `envconfig:"W0Q_LEGAL_ADDRESS" default:"add your address here"`
	LegalMail    string `envconfig:"W0Q_LEGAL_MAIL" default:"add your email address here"`
	LegalPhone   string `envconfig:"W0Q_LEGAL_PHONE" default:"add your phone number here"`
}
