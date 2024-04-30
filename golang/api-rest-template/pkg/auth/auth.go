package auth

import (
	"api-rest/pkg/database"
	"api-rest/pkg/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func LoginHandler(c *gin.Context) {
	var incomingUser models.User
	var dbUser models.User

	//get json body
	if err := c.ShouldBindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := database.DB.Where("username = ?", incomingUser.Username).First(&dbUser).Error; err != nil [
		
	]
}
