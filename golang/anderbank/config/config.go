package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	PortHttp    string
	SecretKey   string
	Issure      string
	TimeExpires string
}

func setenv(fileName string) error {
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
	err := setenv(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env", err)
	}
	return Config{
		Host:        os.Getenv("anderbank-hostdb"),
		Port:        os.Getenv("anderbank-portdb"),
		User:        os.Getenv("anderbank-userdb"),
		Password:    os.Getenv("anderbank-passworddb"),
		DBName:      os.Getenv("anderbank-dbname"),
		PortHttp:    os.Getenv("anderbank-porthttp"),
		SecretKey:   os.Getenv("anderbank-secretkey"),
		Issure:      os.Getenv("anderbank-issure"),
		TimeExpires: os.Getenv("anderbank-timeexpires"),
	}
}
