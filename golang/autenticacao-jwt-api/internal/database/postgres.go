package database

import (
	"aut-jwt/cmd/config"
	"aut-jwt/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDatabase() {
	host := config.LoadEnv().DB_host
	port := config.LoadEnv().DB_port
	user := config.LoadEnv().DB_user
	pass := config.LoadEnv().DB_password
	db := config.LoadEnv().DB_name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Login{}, &models.User{})
}
