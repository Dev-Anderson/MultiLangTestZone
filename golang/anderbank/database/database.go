package database

import (
	"anderbank/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	e := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", e.Host, e.Port, e.User, e.Password, e.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, err
}
