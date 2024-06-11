package router

import (
	"api-dindin/internal/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/user", handler.GetUserHandler) //ok

		v1.POST("/user", handler.CreateUserHadler) //ok

		v1.PUT("/user", handler.EditUserHandler) //ok

		v1.DELETE("/user", handler.DeleteUserHandler)

		v1.GET("/users", handler.GetAllUserHandler) //ok

		v1.GET("/transaction", handler.GetTransactionHandler) //ok

		v1.POST("/transaction", handler.CreateTransactionHandler) //ok

		v1.PUT("/transaction", handler.EditTransactionHandler) //ok

		v1.DELETE("/transaction", handler.DeleteTransationHandler) //ok

		v1.GET("/transactions", handler.GetAllTransactionHandler) //ok

		v1.GET("/cards", handler.GetAllCardHandler) //ok

		v1.GET("/card", handler.GetCardHandler) //ok

		v1.POST("/card", handler.CreateCardHandler) //ok

		v1.PUT("/card", handler.EditCardHandler) //ok

		v1.DELETE("/card", handler.DeleteCardHandler)

		v1.GET("/banks", handler.GetAllBankHandler)

		v1.GET("/bank", handler.GetBankHandler)

		v1.POST("/bank", handler.CreateBankHandler)

		v1.PUT("/bank", handler.EditBankHandler)

		v1.DELETE("/bank", handler.DeleteBankHandler)
	}
}
