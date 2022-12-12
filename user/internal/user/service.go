package user

import (
	"user.com/internal/crypto"
	"user.com/internal/utils/errors"
)

type UserServiceClient interface {
	CreateUser(user UserRequest) error
}

type userService struct {
	repository UserRepositoryClient
	crypto     crypto.CryptoServiceClient
}

func NewUserService(_repository UserRepositoryClient, _crypto crypto.CryptoServiceClient) UserServiceClient {
	return &userService{
		repository: _repository,
		crypto:     _crypto,
	}
}

func (u *userService) CreateUser(req UserRequest) error {

	if exist := u.repository.ExistEmail(req.Email); exist {
		return &errors.ConflictEmailError{}
	}

	hashPass, err := u.crypto.Encrypt(req.MasterPassword)
	if err != nil {
		return err
	}

	var user User
	user.FillFields(req.Name, req.Email, hashPass)
	err = u.repository.Create(&user)
	if err != nil {
		return err
	}

	return nil
}
