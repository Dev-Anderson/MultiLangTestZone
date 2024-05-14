package database

import (
	"api-test/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// InitDB inicializa a conexao com o banco de dados
func InitDB() *sql.DB {
	c, err := config.LoadEnv()
	if err != nil {
		log.Fatal("Falha ao carregar o arquivo de configuracao: ", err)
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
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
