package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Host string
	Port string
	User     string
	Password string
	Dbname   string
}

func SetupConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}