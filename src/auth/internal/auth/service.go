package auth

import (
	"auth.com/internal/crypto"
	"auth.com/internal/token"
	"auth.com/internal/user"
	"auth.com/internal/utils/errors"
	"time"
)

type ServiceClient interface {
	Login(login *Request) (string, error)
	ValidateToken(auth *TokenRequest) (int64, error)
}

type service struct {
	repository  RepositoryClient
	userService user.ServiceClient
	crypto      crypto.ServiceClient
	token       token.ServiceClient
}

func NewAuthService(
	_repository RepositoryClient,
	_userService user.ServiceClient,
	_crypto crypto.ServiceClient,
	_token token.ServiceClient,
) ServiceClient {
	return &service{
		repository:  _repository,
		userService: _userService,
		crypto:      _crypto,
		token:       _token,
	}
}

func (a *service) Login(login *Request) (string, error) {
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

	authModel, err := a.repository.FindAuthByUser(user.Id)
	if err != nil {
		return "", err
	}

	jwt, err := a.token.CreateTokenByID(user.Id)
	if err != nil {
		return "", err
	}

	nowTime := time.Now()
	if authModel == nil {
		authModel, err = a.repository.CreateUserAuth(&Auth{
			UserId:    user.Id,
			CreatedAt: nowTime,
		})
	}

	authModel.Token = jwt
	authModel.ExpiredAt = nowTime.Add(time.Minute * 30)

	if err = a.repository.UpdateToken(authModel); err != nil {
		return "", err
	}
	return authModel.Token, nil
}

func (a *service) ValidateToken(auth *TokenRequest) (int64, error) {
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
