package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleReceive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	fmt.Printf(" [RECEIVER] Dados processados: ID da Ordem %v - Status: %v\n", data["order_id"], data["status"])

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("recebido"))
}

func main() {
	http.HandleFunc("/receive", handleReceive)
	fmt.Println("  [RECEIVER] Aguardando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
