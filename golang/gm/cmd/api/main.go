package main

import (
	"gm/internal/database"
	"gm/pkg/routes"
	"gm/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDBGorm()
	database.ConnectPostgres(utils.ConfigDSN())
}

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(os.Getenv("gm-porthttp"))
}
