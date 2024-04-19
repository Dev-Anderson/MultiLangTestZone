package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("your-secret-key")

// Claims e uma estrutura que representa as informacoes armazenadas no Token JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken gera um novo token JWT par ao usuario especifico
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // expira em 24 horas

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Token não fornecido")
	}

	tokenString = tokenString[7:] // Remova o prefixo "Bearer "

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return fiber.NewError(fiber.StatusUnauthorized, "Token de assinatura inválida")
		}
		return fiber.NewError(fiber.StatusUnauthorized, "Token inválido")
	}
	if !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Token inválido")
	}

	return nil
}
