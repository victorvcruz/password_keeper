package auth

import (
	"auth.com/internal/platform/database"
)

type AuthRepositoryClient interface {
	CreateUserAuth(auth *Auth) (*Auth, error)
	CreateService(auth *Service) (*Service, error)
	CreateServiceAuth(auth *AuthService) (*AuthService, error)
	FindByToken(token string) (*Auth, error)
	FindServiceByToken(service, serviceConn int64, token string) (*AuthService, error)
	FindServiceByName(service string) (*Service, error)
	FindAuthByUser(user int64) (*Auth, error)
	FindAuthByServices(service, serviceConn int64) (*AuthService, error)
	UpdateToken(auth *Auth) error
	UpdateServiceToken(auth *AuthService) error
	UpdateServiceApiToken(id int64, token string) error
}

type userRepository struct {
	db database.DatabaseClient
}

func NewAuthRepository(_db database.DatabaseClient) AuthRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u *userRepository) CreateUserAuth(auth *Auth) (*Auth, error) {
	err := u.db.DB().Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) CreateService(auth *Service) (*Service, error) {
	err := u.db.DB().Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) CreateServiceAuth(auth *AuthService) (*AuthService, error) {
	err := u.db.DB().Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) UpdateToken(auth *Auth) error {
	err := u.db.DB().Model(auth).Where(&Auth{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateServiceToken(auth *AuthService) error {
	err := u.db.DB().Model(auth).Where(&AuthService{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateServiceApiToken(id int64, token string) error {
	err := u.db.DB().Model(&Service{}).Where(&Service{Id: id}).UpdateColumn("api_token", token).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindByToken(token string) (*Auth, error) {
	var auth Auth
	err := u.db.DB().Where(&Auth{Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (u *userRepository) FindServiceByToken(service, serviceConn int64, token string) (*AuthService, error) {
	var auth AuthService
	err := u.db.DB().Where(&AuthService{Service: service, ServiceConn: serviceConn, Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (u *userRepository) FindServiceByName(service string) (*Service, error) {
	var auth Service
	err := u.db.DB().Where(&Service{Name: service}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (u *userRepository) FindAuthByServices(service, serviceConn int64) (*AuthService, error) {
	var auth AuthService
	err := u.db.DB().Where(&AuthService{Service: service, ServiceConn: serviceConn}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (u *userRepository) FindAuthByUser(user int64) (*Auth, error) {
	var auth Auth
	err := u.db.DB().Where(&Auth{UserId: user}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}
