package main

import (
	"fmt"
	"project-filmes-desafio/internal/database"
)

func main() {
	fmt.Println("Conectando com o banco de dados...")
	database.ConnectDatabase()
}
