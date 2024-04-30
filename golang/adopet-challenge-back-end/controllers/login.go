package controllers

import (
	"adopet/database"
	"adopet/models"
	"adopet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel enviar o body" + err.Error(),
		})
		return
	}

	var user models.User
	dbError := database.DB.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Falha ao carregar o usuario" + dbError.Error(),
		})
		return
	}

	if user.Password != services.Sha256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar o token" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
