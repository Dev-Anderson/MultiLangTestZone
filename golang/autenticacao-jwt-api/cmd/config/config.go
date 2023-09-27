package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	SecretKey   string
	Issure      string
	TimeExpire  string
	DB_host     string
	DB_port     string
	DB_user     string
	DB_password string
	DB_name     string
}

func LoadEnv() Env {
	godotenv.Load()
	return Env{
		SecretKey:   os.Getenv("secretKey"),
		Issure:      os.Getenv("issure"),
		TimeExpire:  os.Getenv("timeExpire"),
		DB_host:     os.Getenv("db_host"),
		DB_port:     os.Getenv("db_port"),
		DB_user:     os.Getenv("db_user"),
		DB_password: os.Getenv("db_password"),
		DB_name:     os.Getenv("db_name"),
	}
}
