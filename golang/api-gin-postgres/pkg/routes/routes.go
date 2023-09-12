package routes

import (
	"api-gin-postgres/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.GET("/produto", controllers.GetAllProdutos)

	r.Run(":8081")
}
