package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func LoadEnd() Env {
	godotenv.Load()
	return Env{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
