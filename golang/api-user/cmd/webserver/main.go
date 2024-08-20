package main

import (
	"api-ser/config/env"
	"api-ser/config/logger"
	"log/slog"
)

func main() {
	logger.InitLogger()
	slog.Info("Start API")

	//carregando as congis
	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("Falha ao carregar as variaveis", err, slog.String("pacakge", "main"))
		return
	}
}
