package database

import (
	"database/sql"
	"fmt"
	"project-filmes-desafio/cmd/config"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	c, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.User, c.Password, c.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Falha ao realizar o ping no banco de dadso", err)
	}

	return db, nil
}
