package config

import (
	"bufio"
	"os"
	"strings"
)

func SetEnv(fileName string) error {
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

type ConfigEnv struct {
	Port string
	DSN  string
}

// func LoadEnv() ConfigEnv {
// 	err := setEnv(".env")
// 	if err != nil {
// 		log.Fatal("Erro set env", err)
// 	}
// 	return ConfigEnv{
// 		Port: os.Getenv("api-coffe-port"),
// 		DSN:  os.Getenv("api-coffe-dsn"),
// 	}
// }
