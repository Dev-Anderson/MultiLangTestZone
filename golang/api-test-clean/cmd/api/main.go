package main

import (
	"api-test/cmd/handler"
	"api-test/internal/database"
	"api-test/internal/repository"
	"api-test/internal/routes"
	"api-test/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := database.InitDB()
	defer dbConn.Close()

	// inicializar o serviceo de user com o reposotyr que utilzia conexao com o banco
	userService := &services.UserService{UserRepository: &repository.UserRepository{DB: dbConn}}
	//Inicializar o manipualdro de usuario com o servico de usuario
	userHandler := &handler.UserHandler{UserService: userService}
	//configurar o servidor
	router := gin.Default()
	//config routes
	routes.SetupRoutes(router, userHandler)
	//Execute o servidro
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao rodar o servidor: ", err)
	}
}
