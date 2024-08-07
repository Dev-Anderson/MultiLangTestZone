package db

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=teste sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Up()

	//exemplo de codigo para fazer o down
	// driver, err := postgres.WithInstance(conn, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://./migrations",
	// 	"postgres", driver)
	// if err != nil {
	// 	panic(err)
	// }
	// m.Down()

	return conn, err
}
