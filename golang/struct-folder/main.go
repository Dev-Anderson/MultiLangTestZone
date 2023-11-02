package main

import (
	"api-simples/config"
	"api-simples/database"
	"api-simples/routes"
	"fmt"
)

func main() {
	fmt.Println("Estruturacao de pasta para uma API simples ")
	e, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}
	fmt.Println("Porta da API ", e.PortHttp)
	database.ConnectDatabase()
	routes.Initialize()
}
