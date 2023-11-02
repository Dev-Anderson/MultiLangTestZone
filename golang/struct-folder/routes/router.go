package routes

import (
	"api-simples/config"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	ConfigRoutes(router)

	e, err := config.LoadEnv()
	if err != nil {
		log.Panicln("Falha ao carregar o env", err)
	}
	router.Run("localhost:" + e.PortHttp)
}
