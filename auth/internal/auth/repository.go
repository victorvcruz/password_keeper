package auth

import "gorm.io/gorm"

type AuthRepositoryClient interface {
	CreateUserAuth(auth *Auth) (*Auth, error)
	CreateService(auth *Service) (*Service, error)
	CreateServiceAuth(auth *AuthApi) (*AuthApi, error)
	FindByToken(token string) (*Auth, error)
	FindServiceByToken(service, serviceConn int64, token string) (*AuthApi, error)
	FindServiceByName(service string) (*Service, error)
	FindAuthByUser(user int64) (*Auth, error)
	FindAuthByServices(service, serviceConn int64) (*AuthApi, error)
	UpdateToken(auth *Auth) error
	UpdateServiceToken(auth *AuthApi) error
	UpdateServiceApiToken(id int64, token string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewAuthRepository(_db *gorm.DB) AuthRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u *userRepository) CreateUserAuth(auth *Auth) (*Auth, error) {
	err := u.db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) CreateService(auth *Service) (*Service, error) {
	err := u.db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) CreateServiceAuth(auth *AuthApi) (*AuthApi, error) {
	err := u.db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (u *userRepository) UpdateToken(auth *Auth) error {
	err := u.db.Model(auth).Where(&Auth{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateServiceToken(auth *AuthApi) error {
	err := u.db.Model(auth).Where(&AuthApi{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateServiceApiToken(id int64, token string) error {
	err := u.db.Model(&Service{}).Where(&Service{Id: id}).UpdateColumn("api_token", token).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindByToken(token string) (*Auth, error) {
	var auth Auth
	err := u.db.Where(&Auth{Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (u *userRepository) FindServiceByToken(service, serviceConn int64, token string) (*AuthApi, error) {
	var auth AuthApi
	err := u.db.Where(&AuthApi{Service: service, ServiceConn: serviceConn, Token: token}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (u *userRepository) FindServiceByName(service string) (*Service, error) {
	var auth Service
	err := u.db.Where(&Service{Name: service}).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (u *userRepository) FindAuthByServices(service, serviceConn int64) (*AuthApi, error) {
	var auth AuthApi
	err := u.db.Where(&AuthApi{Service: service, ServiceConn: serviceConn}).Find(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (u *userRepository) FindAuthByUser(user int64) (*Auth, error) {
	var auth Auth
	err := u.db.Where(&Auth{UserId: user}).Find(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}
