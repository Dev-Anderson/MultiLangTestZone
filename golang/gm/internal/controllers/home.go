package controllers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	msg := map[string]string{"message": "API rodando com sucesso!"}
	return c.JSON(msg)
}
