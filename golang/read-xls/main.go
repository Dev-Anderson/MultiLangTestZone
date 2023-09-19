package main

import (
	"log"
	"strings"

	xls "github.com/tealeg/xlsx"
)

type DadosEndereco struct {
	ID       string
	Endereco string
}

func main() {
	// Abra o arquivo XLS
	xlFile, err := xls.OpenFile("assets/arquivo.xlsx")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo XLS: %v", err)
	}

	for _, sheet := range xlFile.Sheets {
		// fmt.Printf("Lendo planilha %s\n", sheet.Name)

		firstRow := true

		for _, row := range sheet.Rows {
			if firstRow {
				firstRow = false
				continue
			}

			// if len(row.Cells) < 2 {
			// 	log.Printf("Linha nao contem informacoes suficientes")
			// 	continue
			// }

			cameraID := row.Cells[0].String()
			endereco := row.Cells[1].String()

			endereco = strings.TrimSpace(endereco)

			dadosEndereco := []DadosEndereco{}

			dadosEndereco = append(dadosEndereco, DadosEndereco{ID: cameraID, Endereco: endereco})

			// fmt.Printf("ID da camera: %s, Endereco: %s\n", cameraID, endereco)
		}
	}
}
