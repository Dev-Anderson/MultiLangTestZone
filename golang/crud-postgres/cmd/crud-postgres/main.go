package main

import (
	database "crud/cmd/db"
	"crud/internal/usecase"
	"fmt"
)

func main() {
	fmt.Println("Criando a tabela Login")
	usecase.CreateTable()
	fmt.Println("Criando registro dentro do banco de dados")
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}
	usecase.CreateLogin(db, "email@gmail.com", "123")

	fmt.Println("Consulta registro no banco de dados")
	fmt.Println(usecase.GetAllLogin(db))

	fmt.Println("Deletando um registro")
	usecase.DeleteLogin(db)

	fmt.Println("Listando os registros novamente depois do delete")
	fmt.Println(usecase.GetAllLogin(db))

	fmt.Println("Update no banco de dados")
	usecase.UpdateLogin(db)

	fmt.Println("Listando os registro novamente depois do update")
	fmt.Println(usecase.GetAllLogin(db))

	fmt.Println("Programa finalizado")
}
