package handlers

import (
	"api-cookies/handlers/models"
	"api-cookies/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token  string    `json:"string"`
	Expire time.Time `json:"expires"`
}

type SessionData struct {
	Token  string    `json:"string"`
	UserId uuid.UUID `json:"userId"`
}

func (lac *LocalApiConfig) SignInHandler(c *gin.Context) {
	var usertoAuth models.UserToAuth
	if err := c.ShouldBindJSON(&usertoAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//do validate here
	validationErrors := utils.ValidateUserAuth(usertoAuth)
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors,
		})
		return
	}

	//fetch  the user from database to match
	foundUser, err := lac.DB.FindUserByEmail(c, userToAuth.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No user found with this email or password",
		})
		return
	}
	if foundUser.Password != usertoAuth.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
}
