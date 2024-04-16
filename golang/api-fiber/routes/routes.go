package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Rotas
	app.Get("/", func(c *fiber.Ctx) error {
		msg := map[string]string{"mesasge": "API rodando com o Fiber"}
		return c.JSON(msg)
	})
}
