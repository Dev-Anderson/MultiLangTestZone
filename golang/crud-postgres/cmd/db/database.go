package database

import (
	"crud/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {
	e, err := config.GetEnv()
	if err != nil {
		fmt.Println("Falha ao carregar as variaveis de ambiente", err.Error())
	}
	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.Host, e.Port, e.User, e.Password, e.DBName)

	fmt.Println(stringConnection)

	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
