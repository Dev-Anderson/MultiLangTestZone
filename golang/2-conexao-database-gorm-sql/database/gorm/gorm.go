package gormdb

import (
	"duas-conexao-gorm-sq/config"
	"duas-conexao-gorm-sq/schema"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	c := config.LoadEnv()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.User, c.Password, c.DBName, c.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&schema.Pedido{}, &schema.TotalVendas{})
	return db, err
}

// func Close(db *gorm.DB) {
// 		defer db.
// }
