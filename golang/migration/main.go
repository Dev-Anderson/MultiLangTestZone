package main

import (
	"database/sql"
	"encoding/json"
	"migrate-teste/migrate/db"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	conn, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/users", ListUsers(conn))

	http.ListenAndServe(":8080", r)
}

type User struct {
	ID        sql.NullInt64
	FirstName string
	LastName  string
	CreateAd  time.Time
}

func ListUsers(conn *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(`SELECT * FROM users`)
		if err != nil {
			return
		}

		var users []User

		for rows.Next() {
			var user User

			err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreateAd)
			if err != nil {
				continue
			}

			users = append(users, user)
		}

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(users)
	}
}
