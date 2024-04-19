package database

import (
	"database/sql"
	"duas-conexao-gorm-sq/config"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	c := config.LoadEnv()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.User, c.Password, c.DBName, c.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db, err
}

func Close(db *sql.DB) {
	defer db.Close()
}
