package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// config
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "teste"
)

// InitDB inicializa a conexao com o banco de dados
func InitDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Falha ao dar um ping no banco de dados ", err)
	}
	fmt.Println("Conexao feita com sucesso com o banco de dados")
	return db
}
