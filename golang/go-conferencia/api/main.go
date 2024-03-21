package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "root:1234@tcp(localhost:3310)/goconference?allowNativePasswords=true"

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/users", listUserHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/users/{id}", getUserHandler)
	mux.HandleFunc("/cpu", CPUIntensiveEndpoint)

	go http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":6060", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Teste HOME", http.StatusOK)
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []*user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, &u)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Erro enconding JSON", http.StatusInternalServerError)
		return
	}

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if _, err := db.Exec("INSERT INTO users (name, email) VALUES(?, ?)", u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var u user

	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, "Error enconding JSON", http.StatusInternalServerError)
		return
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func CPUIntensiveEndpoint(w http.ResponseWriter, r *http.Request) {
	n := 42
	result := fib(n)
	w.Write([]byte(strconv.Itoa(result)))
}
