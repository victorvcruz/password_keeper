package auth

import (
	"auth.com/internal/crypto"
	"auth.com/internal/token"
	"auth.com/internal/user"
	"auth.com/internal/utils/errors"
	"time"
)

type AuthServiceClient interface {
	Login(login LoginRequest) (string, error)
	ValidateToken(auth AuthTokenRequest) (int64, error)
}

type authService struct {
	repository  AuthRepositoryClient
	userService user.UserServiceClient
	crypto      crypto.CryptoServiceClient
	token       token.TokenServiceClient
}

func NewAuthService(
	_repository AuthRepositoryClient,
	_userService user.UserServiceClient,
	_crypto crypto.CryptoServiceClient,
	_token token.TokenServiceClient,
) AuthServiceClient {
	return &authService{
		repository:  _repository,
		userService: _userService,
		crypto:      _crypto,
		token:       _token,
	}
}

func (a *authService) Login(login LoginRequest) (string, error) {
	user, err := a.userService.UserByEmail(login.Email)
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

	nowTime := time.Now()
	err = a.repository.Create(&Auth{
		UserId:    user.Id,
		Token:     jwt,
		CreatedAt: nowTime,
		ExpiredAt: nowTime.Add(time.Hour * 1),
	})

	return jwt, nil
}

func (a *authService) ValidateToken(auth AuthTokenRequest) (int64, error) {
	authModel, err := a.repository.FindByToken(auth.AcessToken)
	if err != nil {
		return 0, err
	}

	if authModel == nil {
		return 0, &errors.InvalidTokenError{}
	}

	if authModel.ExpiredAt.Before(time.Now()) {
		return 0, &errors.ExpiredTokenError{}
	}

	user, err := a.userService.UserById(authModel.UserId)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, &errors.NotFoundIdError{}
	}

	return user.Id, nil
}
