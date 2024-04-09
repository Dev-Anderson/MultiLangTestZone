package main

import (
	"anderbank/database"
	"anderbank/database/table"
	"fmt"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados", err)
	}
	table.TableAccount(db)
	table.TableUser(db)
}
