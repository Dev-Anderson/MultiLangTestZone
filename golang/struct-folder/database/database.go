package database

import (
	"api-simples/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	e, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname%s sllmode=disable",
		e.Host, e.Port, e.User, e.Password, e.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
