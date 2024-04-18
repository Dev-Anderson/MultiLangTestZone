package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	database()
	app := fiber.New()
	routes(app)
	app.Listen(":3000")
}

var DB *gorm.DB

func database() {
	dsn := "user=postgres password=1234 dbname=gm port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

type UF struct {
	gorm.Model
	Nome  string `json:"nome"`
	Sigla string `json:"sigla"`
}

func routes(app *fiber.App) {
	app.Get("/ufs", func(c *fiber.Ctx) error {
		var ufs []UF
		if err := DB.Find(&ufs).Error; err != nil {
			return err
		}
		return c.JSON(ufs)
	})
}
