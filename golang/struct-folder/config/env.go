package config

import (
	"api-simples/schemas"
	"bufio"
	"log"
	"os"
	"strings"
)

func setEnv(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func LoadEnv() (schemas.Environment, error) {
	err := setEnv(".env")
	if err != nil {
		log.Fatal("Erro ao setar as variaveis de ambiente", err)
	}
	return schemas.Environment{
		Host:     os.Getenv("api-simples-hostdb"),
		Port:     os.Getenv("api-simples-portdb"),
		User:     os.Getenv("api-simples-userdb"),
		Password: os.Getenv("api-simples-passworddb"),
		DBName:   os.Getenv("api-simples-dbname"),
		PortHttp: os.Getenv("api-simples-porthttp"),
	}, nil
}
