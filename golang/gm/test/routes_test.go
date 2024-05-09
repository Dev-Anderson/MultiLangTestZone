package test

import (
	"gm/pkg/routes"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRouteHome(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("falha ao criar solicitacao de teste: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("falha ao fazer solicitacao de teste: %v", resp.StatusCode)
	}

	expected := `{"message":"API rodando com sucesso"}`
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("erro ao ler o corpo da resposta: %v", err)
	}
	if string(body) != expected {
		t.Errorf("resposta inesperada: %s", body)
	}
}
