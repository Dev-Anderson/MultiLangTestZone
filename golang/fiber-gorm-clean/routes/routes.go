package routes

import (
	"fiber-gorm/auth"
	"fiber-gorm/models"
	"fiber-gorm/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", home)
	app.Post("/login", login)
	app.Get("/protected", auth.VerifyToken, protectedRoute)

	user := app.Group("/users")
	user.Post("/", createUser)
	user.Get("/:id", getUserByID)
	user.Put("/:id", updateUser)
	user.Delete("/:id", deleteUser)
}

func createUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return c.JSON(user)
}

func getUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.ID = uint(id)
	if err := repository.UpdateUser(user); err != nil {
		return err
	}

	return c.JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err := repository.DeleteUser(uint(id)); err != nil {
		return err
	}
	return c.SendString("Usuario deletado com sucesso")
}

func home(c *fiber.Ctx) error {
	msg := map[string]string{"message": "API rodando"}
	return c.JSON(msg)
}

func login(c *fiber.Ctx) error {
	//logica de autenticacao
	token, err := auth.GenerateToken("username")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token})
}

func protectedRoute(c *fiber.Ctx) error {
	return c.SendString("Rota protegida")
}
