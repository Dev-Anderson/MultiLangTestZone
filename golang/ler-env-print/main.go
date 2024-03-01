package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type EnvConfig struct {
	Config map[string]string
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		Config: make(map[string]string),
	}
}

func (ec *EnvConfig) Load(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo  %s: %v", filename, err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		//ignora linhas vazias ou comentadas
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			ec.Config[key] = value
		}
	}

	return nil
}

func main() {
	envFileName := ".env"
	envConfig := NewEnvConfig()

	err := envConfig.Load(envFileName)
	if err != nil {
		log.Fatal(err)
	}

	//imprimi os dados da estrutura
	fmt.Println("Conteudo da estrutura:")
	for key, value := range envConfig.Config {
		fmt.Printf("%s: %s\n", key, value)
	}
}
