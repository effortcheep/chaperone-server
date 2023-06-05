package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SERVERPORT string
	JWTKEY     string

	DBDRIVER   string
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBDATABASE string
}

var ENV Config

func SetupEnv() {

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("you should run this commond `copy .env.example .env` %v", err)
	}
	err = viper.Unmarshal(&ENV)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
