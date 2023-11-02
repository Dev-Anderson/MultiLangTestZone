package main

import (
	"fmt"
	"os"
	"strconv"
)

func generatePossibilities(input string) []string {
	var possibilities []string

	// Encontrar a posição do primeiro caractere '_'
	pos := -1
	for i, char := range input {
		if char == '_' {
			pos = i
			break
		}
	}

	if pos == -1 {
		// Se não houver mais '_', retornar o input original
		possibilities = append(possibilities, input)
	} else {
		// Recursivamente gerar possibilidades com dígitos de 0 a 9
		for i := 0; i <= 9; i++ {
			newInput := input[:pos] + strconv.Itoa(i) + input[pos+1:]
			subPossibilities := generatePossibilities(newInput)
			possibilities = append(possibilities, subPossibilities...)
		}
	}

	return possibilities
}

func main() {
	input := "2729____40"
	possibilities := generatePossibilities(input)

	batchSize := 1000 // Número máximo de números em cada arquivo
	// batchSize := 3
	for i := 0; i < len(possibilities); i += batchSize {
		batch := possibilities[i:min(i+batchSize, len(possibilities))]
		filename := fmt.Sprintf("output_%d.txt", i/batchSize+1)
		err := writeToFile(filename, batch)
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}
		fmt.Printf("Escrito %d números em %s\n", len(batch), filename)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func writeToFile(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
