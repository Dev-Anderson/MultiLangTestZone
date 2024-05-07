package handlers

import (
	"api-cookies/internal/config"

	"github.com/gin-gonic/gin"
)

type LocalApiConfig struct {
	*config.ApiConfig
}

func (lac *LocalApiConfig) handler_Createuser(c *gin.Context) {

}
