package controllers

import (
	"adopet/database"
	"adopet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTutor(c *gin.Context) {
	var t models.Tutor
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&t)
	c.JSON(http.StatusOK, t)
}

func GetAllTutores(c *gin.Context) {
	var t []models.Tutor
	database.DB.Find(&t)
	c.JSON(http.StatusOK, t)
}

func AlterTutor(c *gin.Context) {
	var t models.Tutor
	id := c.Params.ByName("id")
	database.DB.First(&t, id)

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&t).UpdateColumns(t)
	c.JSON(http.StatusOK, t)
}

func GetIDTutor(c *gin.Context) {
	var t models.Tutor
	id := c.Params.ByName("id")
	database.DB.First(&t, id)

	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tutor na encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, t)
}

func DeleteTutor(c *gin.Context) {
	var t models.Tutor
	id := c.Params.ByName("id")
	database.DB.First(&t, id)

	if t.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tutor nao encontrado",
		})
		return
	}

	database.DB.Delete(&t, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Tutor deletado com sucesso!",
	})
}
