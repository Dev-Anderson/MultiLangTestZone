package database

import (
	"fiber-gorm/config"
	"fiber-gorm/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	c := config.LoadEnv()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.User, c.Password, c.DBName, c.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
