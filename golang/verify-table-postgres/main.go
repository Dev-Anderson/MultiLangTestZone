package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func tableExists(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)")
	var e bool
	err := db.QueryRow(query, tableName).Scan(&e)
	if err != nil {
		return false, err
	}
	return e, nil
}

func connectDB() (*sql.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=teste sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func main() {
	nameTable := "tabela_teste"
	db, err := connectDB()
	if err != nil {
		fmt.Print("Falha ao conectar com o banco de dados", err)
	}

	exists, err := tableExists(db, nameTable)
	if err != nil {
		panic(err)
	}

	fmt.Println(exists)
}
