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
type authRepository struct {
	db database.Client
}

func NewAuthRepository(_db database.Client) AuthRepositoryClient {
	return &authRepository{
		db: _db,
	}
}

func (a *authRepository) CreateUserAuth(auth *Auth) (*Auth, error) {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *authRepository) CreateService(auth *Service) (*Service, error) {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *authRepository) CreateServiceAuth(auth *AuthService) (*AuthService, error) {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *authRepository) UpdateToken(auth *Auth) error {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Model(auth).Where(&Auth{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepository) UpdateServiceToken(auth *AuthService) error {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Model(auth).Where(&AuthService{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepository) UpdateServiceApiToken(id int64, token string) error {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Model(&Service{}).Where(&Service{Id: id}).UpdateColumn("api_token", token).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepository) FindByToken(token string) (*Auth, error) {
	var auth Auth
	err := a.db.DB().Where(&Auth{Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (a *authRepository) FindServiceByToken(service, serviceConn int64, token string) (*AuthService, error) {
	var auth AuthService
	err := a.db.DB().Where(&AuthService{Service: service, ServiceConn: serviceConn, Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (a *authRepository) FindServiceByName(service string) (*Service, error) {
	var auth Service
	err := a.db.DB().Where(&Service{Name: service}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (a *authRepository) FindAuthByServices(service, serviceConn int64) (*AuthService, error) {
	var auth AuthService
	err := a.db.DB().Where(&AuthService{Service: service, ServiceConn: serviceConn}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}

func (a *authRepository) FindAuthByUser(user int64) (*Auth, error) {
	var auth Auth
	err := a.db.DB().Where(&Auth{UserId: user}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	if auth.Id == 0 {
		return nil, nil
	}

	return &auth, nil
}
