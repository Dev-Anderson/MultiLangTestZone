package database

import (
	"anderbank/config"
	"anderbank/schema"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	e := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", e.Host, e.Port, e.User, e.Password, e.DBName)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&schema.User{}, &schema.Login{})
}
