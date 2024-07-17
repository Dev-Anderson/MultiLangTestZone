package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	fmt.Println("Acessando o FTP")
	conn, err := ftp.Dial("url:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Login("user", "pass")
	if err != nil {
		log.Fatal(err)
	}

	diretorio := ""
	nomeArquivo := ""

	entries, err := conn.List(diretorio)
	if err != nil {
		log.Fatal(err)
	}

	var foundFile *ftp.Entry
	for _, entry := range entries {
		if entry.Name == nomeArquivo {
			foundFile = entry
			break
		}
	}

	if foundFile != nil {
		fmt.Printf("Arquivo encontrado: %s\n", foundFile.Name)
		fmt.Printf("Data e hora de modificao: %s\n", foundFile.Time)
	} else {
		fmt.Printf("Arquivo %s nao encontrado\n", nomeArquivo)
	}

	err = conn.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
