package newlogger

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// representa os diferentes niveis de log
type LogLevel int

// constante que recebe o tipo de log do tipo inteiro
const (
	Error LogLevel = iota
	Warning
	Info
)

// Estrutura que contem o nivel atual de log
type Logger struct {
	level LogLevel
}

// Cria uma nova instancia de logger com o nivel padrao
func NewLogger() *Logger {
	levelStr := strings.ToLower(os.Getenv("LOGSLEVEL"))
	level := LogLevelFromString(levelStr)
	return &Logger{level}
}

// convert uma string para o tipo de LogLevel
func LogLevelFromString(levelStr string) LogLevel {
	switch levelStr {
	case "error":
		return Error
	case "warning":
		return Warning
	case "info":
		return Info
	default:
		return Info
	}
}

func levelString(level LogLevel) string {
	switch level {
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Info:
		return "INFO"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) GenerateLog(level LogLevel, code, message string) {
	if level <= l.level {
		typeDelete := os.Getenv("LOGSTYPEDELETE")
		typeDeleteValue, _ := strconv.Atoi(os.Getenv("LOGSTYPEDELETEVALUE"))

		if typeDelete == "DAY" {
			deleteFileInForDay(typeDeleteValue)
		} else if typeDelete == "SIZE" {
			deleteFileInfoMB(typeDeleteValue)
		} else {
			log.Println("invalid delete type")
		}

		os.Mkdir(fmt.Sprintf("%s/%s", path(), "logs"), 0755)
		_, err := os.Stat(nameFileWithPath())
		if err != nil {
			file, err := os.Create(nameFileWithPath())
			if err != nil {
				log.Println("Erro in create file", err)
			}
			defer file.Close()
		}
		logFile, err := os.OpenFile(nameFileWithPath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Erro open file", err)
		}

		logger := log.New(logFile, "", log.LstdFlags)

		stringLog := fmt.Sprintf("[%s] %s - %s", levelString(level), code, message)
		logger.Println(stringLog)
		//se o nivel for info, adicione as informacoes de CPU e memoria
		if level == Info {
			logger.Println(getMemory())
			logger.Println(getCPU())
		}

		defer logFile.Close()
	}
}
