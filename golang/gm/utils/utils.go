package utils

import (
	"fmt"
	"gm/cmd/config"
)

func ConfigDSN() string {
	c, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.User, c.Password, c.DBName, c.Port)
	return dsn
}
