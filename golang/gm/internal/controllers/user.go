package controllers

import (
	"log"
	"strconv"

	"gm/internal/repository"

	"gm/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func Createuser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Printf("Erro ao fazer o parse JSON: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao fazer o parse JSON"})
	}
	if err := repository.CreateUser(user); err != nil {
		log.Printf("Erro ao inserir o usuario err: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao criar o usuario"})
	}
	return c.JSON(user)
}

func GetAlluser(c *fiber.Ctx) error {
	user, err := repository.GetAllUser()
	if err != nil {
		log.Printf("Erro ao buscar usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao buscar usuarios"})
	}

	return c.JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("ID: %d com erro: %s", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao converter o usuario"})
	}

	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		log.Printf("Falha ao buscar o usuario com o ID: %d", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "falha ao buscar o usuario"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("ID: %d com erro: %s", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao converter o ID do usuario"})
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Printf("Erro ao fazer o parse JSON do user: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao fazer o parse JSON do user"})
	}
	user.ID = uint(id)
	if err := repository.UpdateUser(user); err != nil {
		log.Printf("ID: %d com erro ao fazer alteracao erro: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao fazer o update no usuario"})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("ID: %d com erro ao fazer a conversao : ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao converter o ID"})
	}

	if err := repository.DeleteUser(uint(id)); err != nil {
		log.Printf("Falha ao fazer Delete no user com o ID: %d erro: %s", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao fazer o delete do usuario"})
	}

	return c.SendString("Usuario deletado com sucesso!")
}
