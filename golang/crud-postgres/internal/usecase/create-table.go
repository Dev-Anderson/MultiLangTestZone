package usecase

import (
	database "crud/cmd/db"
	"database/sql"
	"fmt"
)

func verifyTable(db *sql.DB, nameTabela string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)")
	var exists bool
	err := db.QueryRow(query, nameTabela).Scan(&exists)
	if err != nil {
		return false, nil
	}
	return exists, nil
}

func CreateTable() {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tableExists, err := verifyTable(db, "login")
	if tableExists == false {
		createTable := `
			CREATE TABLE login (
				ID SERIAL PRIMARY KEY, 
				EMAIL VARCHAR(255), 
				PASSWORD VARCHAR(255)
			)`
		_, err := db.Exec(createTable)
		if err != nil {
			fmt.Println("Erro ao criar a tabela", err.Error())
		}
		fmt.Println("Tabela criada com sucesso!")
	} else {
		fmt.Println("A tabela ja existe")
	}
}
