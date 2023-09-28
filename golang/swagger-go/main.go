package main

import (
	"net/http"
	_ "swagger-go/docs"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Minha API Swagger
// @version 1.0
// @description Exemplo de API Swagger em Go
// @host localhost:8080
// @BasePath /api
func main() {
	r := gin.Default()

	// Rota Swagger
	r.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	// Rota da sua API
	r.GET("/api/hello", helloHandler)

	r.Run(":8080")
}

// @Summary Retorna uma saudação
// @Description Retorna uma saudação simples
// @Produce json
// @Success 200 {string} string "Hello, world!"
// @Router /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
}
