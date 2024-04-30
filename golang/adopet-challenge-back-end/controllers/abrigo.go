package controllers

import (
	"adopet/database"
	"adopet/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAbrigo(c *gin.Context) {
	var a models.Abrigo
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("o que eu estou recebendo:", a)

	database.DB.Create(&a)
	c.JSON(http.StatusOK, a)
}

func GetAllAbrigo(c *gin.Context) {
	var a []models.Abrigo
	database.DB.Find(&a)
	c.JSON(http.StatusOK, a)
}

func AlterAbrigo(c *gin.Context) {
	var a models.Abrigo
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&a).UpdateColumns(a)
	c.JSON(http.StatusOK, a)
}

func GetIDAbrigo(c *gin.Context) {
	var a models.Abrigo
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Abrigo nao encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, a)
}

func DeleteAbrigo(c *gin.Context) {
	var a models.Abrigo
	id := c.Params.ByName("id")
	database.DB.First(&a, id)

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Abrigo nao encontrado",
		})
		return
	}

	database.DB.Delete(&a, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Abriro deletado com sucesso!",
	})
}
