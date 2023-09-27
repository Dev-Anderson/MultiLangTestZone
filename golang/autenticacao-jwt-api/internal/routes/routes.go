package routes

import (
	"aut-jwt/internal/controllers"
	"aut-jwt/internal/middleware"

	"github.com/gin-gonic/gin"
)

func HandlerRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAlluser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		home := main.Group("home", middleware.Auth())
		{
			home.GET("/", controllers.Home)
		}

	}

	return router
}
