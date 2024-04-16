package main

import (
	"fiber-gorm/database"
	"fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
