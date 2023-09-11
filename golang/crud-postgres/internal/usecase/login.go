package usecase

import (
	"database/sql"
	"fmt"
)

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateLogin(db *sql.DB, email, password string) error {
	_, err := db.Exec("insert into login (email, password) values($1, $2)", email, password)
	return err
}

func GetAllLogin(db *sql.DB) ([]Login, error) {
	query := "select id, email, password from login order by id"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao fazer a consulta no banco de dados", err.Error())
	}

	defer rows.Close()

	var l []Login
	for rows.Next() {
		var id int
		var email string
		var password string

		err := rows.Scan(&id, &email, &password)
		if err != nil {
			fmt.Println("Erro ao fazer consulta", err.Error())
		}

		l = append(l, Login{ID: id, Email: email, Password: password})
	}

	return l, nil
}

func DeleteLogin(db *sql.DB) {
	var id int
	db.QueryRow("select id from login limit 1").Scan(&id)
	fmt.Println("ID para deletar = ", id)
	db.Exec("delete from login where id = $1", id)
}

func UpdateLogin(db *sql.DB) {
	var id int
	db.QueryRow("select max(id) from login").Scan(&id)
	fmt.Println("Buscando o ultimo registro da tabela", id)
	db.Exec(`
		update login 
		set email = 'update@gmail.com'
		where id = $1
	`, id)

}
