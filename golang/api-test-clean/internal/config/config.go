package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	PortHTTP string
}

func setEnv(filename string) error {
	file, err := os.Open(filename)
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

func LoadEnv() (Config, error) {
	err := setEnv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo de configuracao: ", err)
	}
	return Config{
		Host:     os.Getenv("api-test-host"),
		Port:     os.Getenv("api-test-port"),
		User:     os.Getenv("api-test-user"),
		Password: os.Getenv("api-test-password"),
		DBName:   os.Getenv("api-test-dbname"),
		PortHTTP: os.Getenv("api-test-porthttp"),
	}, nil
}
