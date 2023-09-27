package services

import (
	"aut-jwt/cmd/config"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: config.LoadEnv().SecretKey,
		issure:    config.LoadEnv().Issure,
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	timeExpires, _ := strconv.ParseInt(config.LoadEnv().TimeExpire, 10, 64)
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(timeExpires)).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", nil
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token invalido %s", token)
		}
		return []byte(s.secretKey), nil
	})
	return err == nil
}
