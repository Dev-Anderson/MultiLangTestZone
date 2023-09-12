package controllers

import (
	"api-gin-postgres/cmd/database"
	"api-gin-postgres/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProdutos(c *gin.Context) {
	var p []models.Produto
	database.DB.Find(&p)
	c.JSON(http.StatusOK, p)
}
