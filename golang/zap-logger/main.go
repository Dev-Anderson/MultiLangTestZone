package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Configurar a configuração do logger
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Configurar a saída do logger para um arquivo
	file, err := os.Create("logs.txt")
	if err != nil {
		log.Fatal("Erro ao criar o arquivo de log: ", err)
	}
	defer file.Close()

	// Configurar um logger que escreve no arquivo
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(file),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		),
	)

	// Configurar o logger padrão para escrever no arquivo
	log.SetOutput(logger.Sugar().Writer())

	// Exemplo de log com campos personalizados
	level := zapcore.InfoLevel
	codigoErro := 200
	horario := "2023-11-14T12:00:00Z"
	descricao := "Operação concluída com sucesso"

	logger.Info("Mensagem de log",
		zap.String("level", level.String()),
		zap.Int("codigo_erro", codigoErro),
		zap.String("horario", horario),
		zap.String("descricao", descricao),
	)

	// Exemplo de log de erro
	codigoErro = 500
	descricao = "Erro ao processar a solicitação"

	logger.Error("Erro ao processar a solicitação",
		zap.String("level", zapcore.ErrorLevel.String()),
		zap.Int("codigo_erro", codigoErro),
		zap.String("horario", horario),
		zap.String("descricao", descricao),
	)

	// Importante: Chamar Sync no final para garantir que todos os logs foram escritos
	logger.Sync()
}
