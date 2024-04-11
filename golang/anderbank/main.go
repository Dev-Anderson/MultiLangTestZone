package main

import (
	"anderbank/database"
	"anderbank/hadler/server"
)

func main() {
	database.ConnectDatabase()

	s := server.NewServer()
	s.Run()
}
