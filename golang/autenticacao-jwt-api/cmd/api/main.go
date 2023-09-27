package main

import (
	"aut-jwt/internal/database"
	"aut-jwt/internal/server"
)

func main() {
	database.ConnectionDatabase()
	s := server.NewServer()
	s.Run()
}
