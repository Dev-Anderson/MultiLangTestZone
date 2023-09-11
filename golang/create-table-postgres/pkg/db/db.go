package db

import (
	"create-table-postgres/cmd/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.LoadEnd().Host, config.LoadEnd().User, config.LoadEnd().Password, config.LoadEnd().Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados", err.Error())
	}

	return db, nil
}
