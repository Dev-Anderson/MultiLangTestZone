package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func connectDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=teste sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db, err
}

type Conta struct {
	ID            int       `json:"id"`
	NumberAccount string    `json:"numberaccount"`
	Balance       float64   `json:"balance"`
	DateCreate    time.Time `json:"datecreate"`
	DateUpdate    time.Time `json:"dateupdate"`
	DateDelete    time.Time `json:"datedelete"`
	Debit         bool      `json:"debit"`
	Credit        bool      `json:"credit"`
	Active        bool      `json:"active"`
}

func (c *Conta) Conta() ([]Conta, error) {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `select id, "numberAccount", balance, "dateCreate", "dateUpdate", "dateDelete", debit, credit, active from anderbank_account`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var asdf []*Conta
	for rows.Next() {
		var ct Conta
		err := rows.Scan(
			&ct.ID,
			&ct.NumberAccount,
			&ct.Balance,
			&ct.DateCreate,
			&ct.DateUpdate,
			&ct.DateDelete,
			&ct.Debit,
			&ct.Credit,
			&ct.Active,
		)
		if err != nil {
			return nil, err
		}

		asdf = append(asdf, &ct)
	}
	return asdf, nil
}

func main() {

}
