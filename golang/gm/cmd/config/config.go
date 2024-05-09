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
	PortHttp string
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

func LoadEnv() (Config, error) {
	err := setEnv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo de configuracao", err)
	}
	return Config{
		Host:     os.Getenv("gm-host"),
		Port:     os.Getenv("gm-port"),
		User:     os.Getenv("gm-user"),
		Password: os.Getenv("gm-password"),
		DBName:   os.Getenv("gm-dbname"),
		PortHttp: os.Getenv("gm-porthttp"),
	}, nil
}
