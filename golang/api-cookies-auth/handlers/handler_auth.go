package handlers

import (
	"api-cookies/handlers/models"
	"api-cookies/utils"
	"encoding/json"
	"net/http"
	"os"
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
	var userToAuth models.UserToAuth
	if err := c.ShouldBindJSON(&userToAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// do validation here
	validationErrors := utils.ValidateUserAuth(userToAuth)
	if len(validationErrors) > 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErrors,
		})
		return
	}

	//fetch the user from database to match
	foundUser, err := lac.DB.FindUserByEmail(c, userToAuth.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No user found with this email or password",
		})
		return
	}
	if foundUser.Password != userToAuth.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Email: userToAuth.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// till here you are having token
	sessionID := uuid.New().String()

	//prepare the session data in redis
	sessionData := map[string]interface{}{
		"token":  tokenString,
		"userId": foundUser.ID,
	}

	sessionDataJSON, err := json.Marshal(sessionData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encode the json data",
		})
		return
	}

	// saving the session data in redis
}
