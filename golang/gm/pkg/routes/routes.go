package routes

import (
	"gm/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	users := app.Group("/user")
	users.Post("/", controllers.Createuser) //falta teste
	users.Get("/", controllers.GetAlluser)
}
