package main

import (
	"errors"

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

type Cidade struct {
	gorm.Model
	Nome       string `json:"nome"`
	CodigoIBGE string `json:"codigoibge"`
	UFID       uint   `json:"uf_id"`
	UF         UF     `gorm:"foreignKey:"UFID"`
}

func routes(app *fiber.App) {
	app.Get("/cidades", func(c *fiber.Ctx) error {
		sigla := c.Query("uf_sigla")
		if sigla == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Sigla da UF nao fornecidad"})
		}

		//execute a consulta para obter as cidades com a sigla de UF fornecida
		var cidades []Cidade
		if err := DB.Preload("UF").Joins("UF").Where(`"UF".sigla = ?`, sigla).Find(&cidades).Error; err != nil {
			// if err := DB.Preload("UF").Joins("UF").Where("uf = sigla = ?", sigla).Find(&cidades).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Nenhuma cidade encontrada para esta UF"})
			}
			return err
		}
		// construa a resposta JSON com os campos desejados
		var response []map[string]interface{}
		for _, cidade := range cidades {
			response = append(response, map[string]interface{}{
				"nome":  cidade.Nome,
				"ibge":  cidade.CodigoIBGE,
				"sigla": cidade.UF.Sigla,
			})
		}
		return c.JSON(response)
	})
}
