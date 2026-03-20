package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SaleEvent struct {
	OrderID string  `json:"order_id"`
	Status  string  `json:"status"`
	Total   float64 `json:"total"`
}

func sendWebhook(targetURL string, payload SaleEvent) {
	// 1. converter struct para JSON
	jsonData, _ := json.Marshal(payload)

	// 2. configurar o cliente http com timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	fmt.Printf(" [SENDER] Disparando evento para %s...\n", targetURL)

	//3. enviar o post
	resp, err := client.Post(targetURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf(" [ERRO] Falha ao enviar webhook: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(" [SENDER] Webhook entregue com sucesso!")
	} else {
		fmt.Printf(" [SENDER] Receptor respondeu com erro: %d\n", resp.StatusCode)
	}
}

func main() {
	// dados do evento
	e := SaleEvent{
		OrderID: "abc-123",
		Status:  "teste",
		Total:   120.33,
	}

	// url do nosso receptor (ajuste se necessario)
	target := "http://localhost:8080/receive"

	//simulado o disparo
	sendWebhook(target, e)
}
