package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetEnv() (Env, error) {
	godotenv.Load()
	return Env{
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	}, nil
}
