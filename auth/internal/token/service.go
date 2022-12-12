package token

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type TokenServiceClient interface {
	CreateTokenByID(id int) (string, error)
}

type TokenService struct {
	key string
}

func NewTokenService() TokenServiceClient {
	return &TokenService{
		key: os.Getenv("JWT_TOKEN_KEY"),
	}
}

func (t *TokenService) CreateTokenByID(id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(t.key))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
