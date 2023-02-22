package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	// Postgresql
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	return Config{
		Port: os.Getenv("PORT"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASS"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresDB: os.Getenv("POSTGRES_DB"),
	}
}