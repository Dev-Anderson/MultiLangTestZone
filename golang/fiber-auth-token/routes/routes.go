package routes

import (
	"fiber-auth-jwt/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", home)
	app.Post("/login", login)
	app.Get("/protected", auth.VerifyToken, protectedRoute)
}

func home(c *fiber.Ctx) error {
	msg := map[string]string{"message": "Api de testes com o Fiber e Auth rodando"}
	return c.JSON(msg)
}

func login(c *fiber.Ctx) error {
	//logica da autenticacao
	token, err := auth.GenerateToken("username")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"Token": token})
}

func protectedRoute(c *fiber.Ctx) error {
	msg := map[string]string{"message": "teste"}
	return c.JSON(msg)
}
