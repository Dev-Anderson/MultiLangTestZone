package table

import (
	"database/sql"
	"fmt"
)

func TableAccount(db *sql.DB) {
	tableExists, err := verifyTable(db, "anderbank_account")
	if tableExists == false {
		create := `
		create table "anderbank_account"(
			"id" serial primary key, 
			"numberAccount" text, 
			"balance" numeric, 
			"dateCreate" timestamp, 
			"dateUpdate" timestamp, 
			"dateDelete" timestamp, 
			"debit" bool, 
			"credit" bool, 
			"active" bool);`

		_, err := db.Exec(create)
		if err != nil {
			fmt.Println("Erro ao criar tabela anderbank_account", err.Error())
		}
		fmt.Println("Table anderbank_account criada com sucesso")
	} else {
		fmt.Println("Tablea anderbank_account ja existe")
	}
	if err != nil {
		fmt.Println("Erro ao verificar a tablea anderbank_account", err.Error())
	}
}
