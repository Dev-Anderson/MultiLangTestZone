package routes

import (
	"api-test/cmd/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	//definido rotas
	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.Createuser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	// router.PATCH()
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
