package main

import (
	"log/slog"
	"test-logger/config"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "HIDDEN"),
	)
}

func main() {
	config.InitLogger()

	user := user{
		Name:     "Anderson da Silva",
		Age:      34,
		Password: "1234",
	}

	slog.Info("Iniciando a aplicacao")
	slog.Error("teste de log de erro")
	slog.Debug("teste de log debug")
	slog.Warn("teste de log de warg")
	slog.Info("Criando o usuario", "user", user.LogUser())
}
