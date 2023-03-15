package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

type TokenServiceClient interface {
	CreateTokenByID(id int64) (string, error)
	DecodeTokenReturnId(token string) (string, error)
}

type TokenService struct {
	key string
}

func NewTokenService() TokenServiceClient {
	return &TokenService{
		key: os.Getenv("JWT_TOKEN_KEY"),
	}
}

func (t *TokenService) CreateTokenByID(id int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	tokenString, err := token.SignedString([]byte(t.key))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (t *TokenService) DecodeTokenReturnId(token string) (string, error) {

	tokenStr := strings.ReplaceAll(token, "Bearer ", "")
	tokenDecode := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, tokenDecode, func(tokenStr *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	switch tokenDecode["id"].(type) {
	case string:
		return tokenDecode["id"].(string), nil
	case float64:
		return fmt.Sprintf("%0.f", tokenDecode["id"].(float64)), nil
	case int:
		return fmt.Sprintf("%d", tokenDecode["id"].(int)), nil
	}
	return "", errors.New("unprocessable token")
}
