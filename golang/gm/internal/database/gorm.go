package database

import (
	"fmt"
	"gm/cmd/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBGorm *gorm.DB

func ConnectDBGorm() {
	c, err := config.LoadEnv()
	if err != nil {
		panic("Erro ao carregar arquivo de configuracao" + err.Error())
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.User, c.Password, c.DBName, c.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar com o banco de dados: " + err.Error())
	}

	DBGorm = db
}
