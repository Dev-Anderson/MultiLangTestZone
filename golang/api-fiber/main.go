package main

import (
	"api-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Criando uma nova instancia do app fiber
	app := fiber.New()

	//Iniciando as rotas
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
