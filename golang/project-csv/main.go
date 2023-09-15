package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("arquivo.csv")
	if err != nil {
		log.Fatal("Falha ao abir o arquivo")
	}

	defer file.Close()

	customReader := csv.NewReader(bufio.NewReader(file))
	customReader.Comma = ';'

	header, err := customReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	idIndex := -1
	descricaoIndex := -1

	for i, col := range header {
		if strings.TrimSpace(col) == "ID" {
			idIndex = i
		} else if strings.TrimSpace(col) == "DESCRICACAO" {
			descricaoIndex = i
		}
	}

	if idIndex == -1 || descricaoIndex == -1 {
		fmt.Println("Colunas 'ID' e 'DESCRICACAO' não encontradas no cabeçalho.")
		return
	}

	// Leia as linhas do CSV
	for {
		// Leia uma linha do CSV
		record, err := customReader.Read()

		// Verifique se chegamos ao final do arquivo
		if err != nil {
			break
		}

		// Acesse os campos usando os índices mapeados
		id := record[idIndex]
		descricao := record[descricaoIndex]

		// Faça o que quiser com os dados lidos
		fmt.Printf("ID: %s, Descricao: %s\n", id, descricao)
	}
}
