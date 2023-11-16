package main

import (
	"log/slog"
	"os"
)

func main() {
	//padrao 01 sem JSON
	slog.Info("teste de log")
	slog.Error("teste de erro")
	slog.Warn("Teste de warning")

	//padrao 02 com JSON
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	l.Info("teste de info")
	l.Error("teste de info")
	l.Warn("teste de info")

	//outro teste
	slog.SetDefault(l)
	l.Info("teste do L default")

	l = slog.Default().WithGroup("teste-withGroup")
	l.Info("Logado com sucesso")

}
