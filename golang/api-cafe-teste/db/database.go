package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConnn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxIdleConns(maxOpenDbConnn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Print("Error", err)
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")
	return nil
}
