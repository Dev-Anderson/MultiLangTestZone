package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserPayload struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// 1. validar se o metodo e post (webhook quse sempre usam POST)
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. ler o corpo da requisicao
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// 3. Desserializar o JSON para a nossa struct
	var user UserPayload
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	// 4. logica de negocio
	fmt.Printf(" [webhook recebido] novo usuario? %s (%s) - ID: %d\n", user.Name, user.Email, user.ID)

	// 5. Responder com 200 OK imediatament
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook processado com sucesso"))
}

func main() {
	// definindo a rota do webhook
	http.HandleFunc("/webhooks/new-user", webhookHandler)

	fmt.Println("Servidor webhook rodadno na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
