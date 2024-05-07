package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lac *LocalApiConfig) HandlerReadiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
