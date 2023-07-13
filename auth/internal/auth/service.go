package auth

import (
	"auth.com/internal/crypto"
	"auth.com/internal/token"
	"auth.com/internal/user"
	"auth.com/internal/utils/errors"
	"time"
)

type AuthServiceClient interface {
	Login(login *LoginRequest) (string, error)
	LoginService(login *LoginServiceRequest) (string, error)
	ValidateToken(auth *AuthTokenRequest) (int64, error)
	ValidateServiceToken(auth *AuthTokenService) (int64, error)
	RegisterService(register *Register) (string, error)
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

func (a *authService) Login(login *LoginRequest) (string, error) {
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

func (a *authService) LoginService(login *LoginServiceRequest) (string, error) {
	service, err := a.repository.FindServiceByName(login.Service)
	if err != nil {
		return "", err
	}

	if service == nil {
		return "", &errors.NotFoundServiceError{}
	}

	if service.ApiToken != login.ApiToken {
		return "", &errors.UnauthorizedApiTokenError{}
	}

	serviceConn, err := a.repository.FindServiceByName(login.ServiceConn)
	if err != nil {
		return "", err
	}

	if serviceConn == nil {
		return "", &errors.NotFoundServiceError{}
	}

	authModel, err := a.repository.FindAuthByServices(service.Id, serviceConn.Id)
	if err != nil {
		return "", err
	}

	nowTime := time.Now()
	if authModel == nil {
		authModel, err = a.repository.CreateServiceAuth(&AuthApi{
			Service:     service.Id,
			ServiceConn: serviceConn.Id,
			ApiToken:    service.ApiToken,
			CreatedAt:   nowTime,
			ExpiredAt:   nowTime,
		})
	}

	authModel.Token, err = a.token.CreateTokenByID(authModel.Id)
	if err != nil {
		return "", err
	}

	authModel.ExpiredAt = nowTime.Add(time.Minute * 30)

	if err = a.repository.UpdateServiceToken(authModel); err != nil {
		return "", err
	}

	return authModel.Token, nil
}

func (a *authService) ValidateToken(auth *AuthTokenRequest) (int64, error) {
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

func (a *authService) ValidateServiceToken(auth *AuthTokenService) (int64, error) {
	service, err := a.repository.FindServiceByName(auth.Service)
	if err != nil {
		return 0, err
	}

	if service == nil {
		return 0, &errors.NotFoundServiceError{}
	}

	serviceConn, err := a.repository.FindServiceByName(auth.ServiceConn)
	if err != nil {
		return 0, err
	}

	if serviceConn == nil {
		return 0, &errors.NotFoundServiceError{}
	}

	authModel, err := a.repository.FindServiceByToken(service.Id, serviceConn.Id, auth.AcessToken)
	if err != nil {
		return 0, err
	}

	if authModel == nil {
		return 0, &errors.InvalidTokenError{}
	}

	if authModel.ExpiredAt.Before(time.Now()) {
		return 0, &errors.ExpiredTokenError{}
	}

	return authModel.Id, nil
}

func (a *authService) RegisterService(register *Register) (string, error) {
	service, err := a.repository.FindServiceByName(register.Service)
	if err != nil {
		return "", err
	}

	if service != nil {
		return service.ApiToken, nil
	}

	jwt, err := a.token.CreateRandomToken()
	if err != nil {
		return "", err
	}

	nowTime := time.Now()
	_, err = a.repository.CreateService(&Service{
		Name:      register.Service,
		ApiToken:  jwt,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	})
	if err != nil {
		return "", err
	}

	return jwt, nil
}
