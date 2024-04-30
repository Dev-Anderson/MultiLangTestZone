package routes

import (
	"adopet/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		abrigo := main.Group("abrigo")
		{
			abrigo.POST("/", controllers.CreateAbrigo)
			abrigo.GET("/", controllers.GetAllAbrigo)
			abrigo.GET("/:id", controllers.GetIDAbrigo)
			abrigo.PATCH("/:id", controllers.AlterAbrigo)
			abrigo.DELETE("/:id", controllers.DeleteAbrigo)
		}
		tutor := main.Group("tutor")
		{
			tutor.POST("/", controllers.CreateTutor)
			tutor.GET("/", controllers.GetAllTutores)
			tutor.GET("/:id", controllers.GetIDTutor)
			tutor.PATCH("/:id", controllers.AlterTutor)
			tutor.DELETE("/:id", controllers.DeleteTutor)
		}
	}

	return router
}
