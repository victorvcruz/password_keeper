package auth

import (
	"auth.com/internal/crypto"
	"auth.com/internal/token"
	"auth.com/internal/user"
	"auth.com/internal/utils/errors"
)

type AuthServiceClient interface {
	Login(login LoginRequest) (string, error)
	ValidateToken(auth AuthTokenRequest) (string, error)
}

type authService struct {
	repository user.UserRepositoryClient
	crypto     crypto.CryptoServiceClient
	token      token.TokenServiceClient
}

func NewAuthService(
	_repository user.UserRepositoryClient,
	_crypto crypto.CryptoServiceClient,
	_token token.TokenServiceClient,
) AuthServiceClient {
	return &authService{
		repository: _repository,
		crypto:     _crypto,
		token:      _token,
	}
}

func (a *authService) Login(login LoginRequest) (string, error) {

	user, err := a.repository.UserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", &errors.NotFoundEmailError{}
	}

	decryptPass, err := a.crypto.Decrypt(user.MasterPassword)
	if decryptPass != login.Password {
		return "", &errors.UnauthorizedPasswordError{}
	}

	jwt, err := a.token.CreateTokenByID(user.Id)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (a *authService) ValidateToken(auth AuthTokenRequest) (string, error) {

	id, err := a.token.DecodeTokenReturnId(auth.AcessToken)
	if err != nil {
		return "", err
	}

	user, err := a.repository.UserById(id)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", &errors.NotFoundIdError{}
	}

	return id, nil
}
