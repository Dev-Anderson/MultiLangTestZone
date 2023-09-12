package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	User     string
	Password string
	DBName   string
}

func LoadEnd() Env {
	godotenv.Load()
	return Env{
		Host:     os.Getenv("db_host"),
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	}
}
