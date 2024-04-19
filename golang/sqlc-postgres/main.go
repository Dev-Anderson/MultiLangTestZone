package main

import (
	"context"
	"fmt"
	"sqlc/database"

	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

func connection() (*pgx.Conn, error) {
	dsn := "postgres://postgres:1234@localhost:5432/gm"

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(err)
	}

	return conn, nil
}

func getAllUser() {
	q := run()
	users, err := q.GetAllUsers(ctx)
	if err != nil {
		fmt.Println("Falha ao buscar todos os usuarios...", err.Error())
	}
	fmt.Println("Result users: ", users)
}

func getUser(id int64) {
	q := run()
	user, err := q.GetUser(ctx, id)
	if err != nil {
		fmt.Println("Falha ao buscar usuario", err.Error())
	}
	fmt.Println("Result user: ", user)
}

func insertUser(name, email, password string) {
	q := run()
	insertUser, err := q.InsertUser(ctx, database.InsertUserParams{Name: name, Email: email, Password: password})
	if err != nil {
		fmt.Println("Falha ao inserir registro do usuario", err.Error())
	}
	fmt.Println("InserUser: ", insertUser)
}

func deleteuser(id int64) {
	q := run()
	deleteUser := q.DeleteUser(ctx, id)
	fmt.Println("deleteuser: ", deleteUser)
}

func run() *database.Queries {
	c, _ := connection()
	return database.New(c)

}

func main() {
	getAllUser()
	getUser(2)
	insertUser("Anderson da silva", "andersondasilva@gmail.com", "andersondasilva")
	// deleteuser(1)
}
