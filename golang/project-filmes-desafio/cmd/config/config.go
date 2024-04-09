package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Env struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func setEnv(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ports := strings.SplitN(line, "=", 2)
		if len(ports) == 2 {
			key := ports[0]
			value := ports[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func LoadEnv() (Env, error) {
	err := setEnv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env", err)
	}
	return Env{
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		Database: os.Getenv("db_database"),
	}, nil
}
