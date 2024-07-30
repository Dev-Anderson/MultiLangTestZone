package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Função para obter o nome do arquivo CSV do usuário
func getFileName() string {
	var fileName string
	fmt.Print("Digite o nome do arquivo CSV: ")
	fmt.Scanln(&fileName)
	return fileName
}

// Função para ler o conteúdo do arquivo CSV
func readCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo: %w", err)
	}

	return records, nil
}

// Função para exibir o menu e obter a escolha da categoria do usuário
func getChoice() string {
	fmt.Println("Selecione a categoria para contar tarefas:")
	fmt.Println("1. Gestão de Frotas")
	fmt.Println("2. Gestão de Manutenção")
	fmt.Println("3. Truckpag Bank")
	fmt.Println("4. Voltar/Sair do programa...")
	fmt.Print("Digite o número da sua escolha: ")

	var choice string
	fmt.Scanln(&choice)
	return choice
}

// Função para exibir o menu e obter a operação do usuário
func getOperation() string {
	fmt.Println("Selecione a operação:")
	fmt.Println("1. Total de registros")
	fmt.Println("2. Total de pontos")
	fmt.Println("3. Contar registros por intervalo de datas")
	fmt.Println("4. Voltar")
	fmt.Print("Digite o número da sua escolha: ")

	var operation string
	fmt.Scanln(&operation)
	return operation
}

// Função para contar registros para uma categoria específica
func countCategory(records [][]string, category string) int {
	var count int
	for _, record := range records {
		if len(record) > 0 { // Verifica se o registro não está vazio
			if record[1] == category {
				count++
			}
		}
	}
	return count
}

// Função para somar tamanhos para uma categoria específica
func sumSizeByCategory(records [][]string, category string) int {
	var totalSize int
	for _, record := range records {
		if len(record) > 0 { // Verifica se o registro não está vazio
			if record[1] == category {
				size, err := strconv.Atoi(record[9])
				if err == nil {
					totalSize += size
				}
			}
		}
	}
	return totalSize
}

func exitProgram() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func getdateRange() (time.Time, time.Time, error) {
	var startDateStr, endDateStr string

	fmt.Print("Digite a data de inicio (YYYY-MM-DD): ")
	fmt.Scanln(&startDateStr)
	fmt.Printf("Digite a data final (YYYY-MM-DD): ")
	fmt.Scanln(&endDateStr)

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("data de inicio invalida: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("data de fim invalida: %w", err)
	}

	return startDate, endDate, nil
}

func countRecordsByDateRange(records [][]string, startDate, endDate time.Time) int {
	var count int
	for _, record := range records {
		if len(record) > 0 {
			lastMovedStr := record[10] //indice da coluna last_moved
			lastMoved, err := time.Parse("2006-01-02 15:04:05", lastMovedStr)
			if err == nil && !lastMoved.Before(startDate) && !lastMoved.After(endDate) {
				count++
			}
		}
	}
	return count
}

func main() {
	fileName := getFileName()
	records, err := readCSV(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		choice := getChoice()

		if choice == "4" {
			exitProgram()
		}

		var category string
		switch choice {
		case "1":
			category = "Gestão de Frotas"
		case "2":
			category = "Gestão de Manutenção"
		case "3":
			category = "Truckpag Bank"
		default:
			fmt.Println("Escolha inválida")
			continue
		}

		operation := getOperation()
		switch operation {
		case "1":
			count := countCategory(records, category)
			fmt.Printf("Qtde de tarefas '%s': %d\n", category, count)
		case "2":
			totalSize := sumSizeByCategory(records, category)
			fmt.Printf("Tamanho total de '%s': %d\n", category, totalSize)
		case "3":
			startDate, endDate, err := getdateRange()
			if err != nil {
				fmt.Println(err)
				continue
			}
			count := countRecordsByDateRange(records, startDate, endDate)
			fmt.Printf("Numero de pontos por intervalo de data de %s a %s: %d\n", startDate, endDate, count)
		case "4":
			continue
		default:
			fmt.Println("Operação inválida")
		}
	}
}
