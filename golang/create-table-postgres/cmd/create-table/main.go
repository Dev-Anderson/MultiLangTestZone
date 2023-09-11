package main

import (
	"create-table-postgres/internal/entity"
	"create-table-postgres/pkg/db"
	"fmt"
)

func main() {
	fmt.Println("Conexao com o banco de dados")
	db.ConnectDatabase()

	fmt.Println("Criando a tabela")
	entity.Table()
}
