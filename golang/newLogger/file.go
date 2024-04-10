package newlogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func path() string {
	p, err := os.Getwd()
	if err != nil {
		return ""
	}
	return p
}

func pathtoFile() string {
	return fmt.Sprintf("%s/%s", path(), "logs")
}

func nameFile() string {
	return fmt.Sprintf("%s.log", currentDate())
}

func nameFileWithPath() string {
	return fmt.Sprintf(pathtoFile() + "/" + nameFile())
}

// funcao para deletar o arquivo por dia, repassar apenas o valor dia como inteiro
func deleteFileInForDay(day int) {
	folder, err := os.Open(pathtoFile())
	if err != nil {
		log.Println("Error when searchind for file", err)
	}
	defer folder.Close()

	files, err := folder.Readdir(-1)
	if err != nil {
		log.Println("Error reading the directory", err)
	}

	for _, file := range files {
		if now().Sub(file.ModTime()) > (time.Hour * 24 * time.Duration(day)) {
			os.Remove(pathtoFile() + "/" + file.Name())
		}
	}
}

func getFileSize() (int, error) {
	info, err := os.Stat(nameFileWithPath())
	if err != nil {
		return 0, err
	}

	sizeMB := int((info.Size() + 512*1024) / (1024 * 1024))
	return sizeMB, nil
}

// funcao para deletar o arquivo pelo tamanho, repassar apenas o valo como inteiro
func deleteFileInfoMB(sizeLimite int) {
	sizeCurrent, err := getFileSize()
	if err != nil {
		log.Println("Erro in get size file current", err)
	}

	if sizeCurrent >= sizeLimite {
		err := os.Remove(nameFileWithPath())
		if err != nil {
			log.Println("Error in remove file log", err)
		}
	}
}
