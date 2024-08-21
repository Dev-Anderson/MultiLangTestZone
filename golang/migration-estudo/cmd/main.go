package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	postgresURI := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Panic(err)
	}

	fmt.Println("Conectado com o banco de dados")

	select {}
}
