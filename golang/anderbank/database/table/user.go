package table

import (
	"database/sql"
	"fmt"
)

func TableUser(db *sql.DB) {
	tableExists, err := verifyTable(db, "anderbank_user")
	if tableExists == false {
		create := `
		CREATE TABLE "anderbank_user" (
			"id" serial primary key, 
			"name" text, 
			"email" text, 
			"password" text, 
			"dateCreate" timestamp, 
			"dateAtualization" timestamp, 
			"dateDelete" timestamp, 
			"active" bool 
		);`

		_, err := db.Exec(create)
		if err != nil {
			fmt.Println("Erro ao criar table anderbank_user", err.Error())
		}
		fmt.Println("Tabela anderbank_user criada com sucesso")
	} else {
		fmt.Println("table anderbank_user ja existe")
	}
	if err != nil {
		fmt.Println("Erro ao verificar a tabela anderbank_user", err.Error())
	}
}
