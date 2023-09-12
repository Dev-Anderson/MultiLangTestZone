package database

import (
	"api-gin-postgres/config"
	"api-gin-postgres/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.LoadEnd().Host, config.LoadEnd().User, config.LoadEnd().Password, config.LoadEnd().DBName)

	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados", err.Error())
	}

	DB.AutoMigrate(&models.Produto{})
}
