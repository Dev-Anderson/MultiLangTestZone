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

func LoadEnv() Config {
	err := setEnv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo de configuracao ", err)
	}
	return Config{
		Host:     os.Getenv("anderson-host"),
		Port:     os.Getenv("anderson-port"),
		User:     os.Getenv("anderson-user"),
		Password: os.Getenv("anderson-password"),
		DBName:   os.Getenv("anderson-dbname"),
	}
}
