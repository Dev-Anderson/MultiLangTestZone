package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("fiber-teste-auth")

// Claims e uma estrutura que representa as informacoes armazenadas no Token JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Gera um novo token jWT para o usuario especifico, repassando o campo "username"
func GenerateToken(username string) (string, error) {
	expiratinTime := time.Now().Add(24 * time.Hour) // expira em 24 horas

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiratinTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func VerifyToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Token nao fornecido")
	}

	tokenString = tokenString[7:] // removendo o campo "Bearer "

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return fiber.NewError(fiber.StatusUnauthorized, "Token de assinatura invalida")
		}
		return fiber.NewError(fiber.StatusUnauthorized, "Token invalido")
	}
	if !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Token invalido")
	}
	return nil
}
