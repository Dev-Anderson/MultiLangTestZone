package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Configs struct {
	Host        string
	Port        string
	User        string
	Password    string
	Database    string
	SecretKey   string
	Issure      string
	TimeExpires string
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

func LoadEnv() (Configs, error) {
	err := setEnv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo de configuracao")
	}
	return Configs{
		Host:        os.Getenv("ADOPET_HOST"),
		Port:        os.Getenv("ADOPET_PORT"),
		User:        os.Getenv("ADOPET_USER"),
		Password:    os.Getenv("ADOPET_PASSWORD"),
		Database:    os.Getenv("ADOPET_DATABASE"),
		SecretKey:   os.Getenv("ADOPET_SECRETKEY"),
		Issure:      os.Getenv("ADOPET_ISSURE"),
		TimeExpires: os.Getenv("ADOPET_TIMEEXPIRES"),
	}, nil
}
