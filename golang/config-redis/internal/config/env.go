package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	Port     string
	Password string
}

func SetEnv() Env {
	godotenv.Load()
	return Env{
		Host:     os.Getenv("REDIS-HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS-PASSWORD"),
	}
}
