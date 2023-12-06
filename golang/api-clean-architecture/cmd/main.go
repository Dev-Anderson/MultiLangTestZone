package main

import (
	"api-clean-archetecture/pkg/delivery/http"
	"api-clean-archetecture/pkg/entity"
	"api-clean-archetecture/pkg/repository"
	"api-clean-archetecture/pkg/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "string de conexao")
	if err != nil {
		panic("Failed to connect to databse")
	}

	db.AutoMigrate(&entity.User{})

	userRepository := repository.NewUserReposity(db)
	userUseCase := &usecase.UserUseCase{UserReposity: userRepository}
	userHandler := &http.UserHandler{UserUseCase: *userUseCase}

	// Configurar o roteador Gin
	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

}
