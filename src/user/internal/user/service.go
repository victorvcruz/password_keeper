package user

import (
	"user.com/internal/crypto"
	"user.com/internal/utils/errors"
)

type ServiceClient interface {
	CreateUser(user Request) error
	UpdateUser(id int64, req Request) error
	FindUserById(id int64) (*User, error)
	FindUserByData(id, name, email string) (*User, error)
	DeleteUser(id int64) error
}

type service struct {
	ServiceClient
	repository RepositoryClient
	crypto     crypto.ServiceClient
}

func NewUserService(_repository RepositoryClient, _crypto crypto.ServiceClient) ServiceClient {
	return &service{
		repository: _repository,
		crypto:     _crypto,
	}
}

func (u *service) CreateUser(req Request) error {

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

func (u *service) UpdateUser(id int64, req Request) error {

	if exist := u.repository.ExistEmail(req.Email); exist {
		return &errors.ConflictEmailError{}
	}

	user, err := u.repository.FindById(id)
	if err != nil {
		return err
	}

	var hashPass string
	if req.MasterPassword == "" {
		hashPass, err = u.crypto.Encrypt(req.MasterPassword)
		if err != nil {
			return err
		}
	}

	user.ToUpdate(req.Name, req.Email, hashPass)
	err = u.repository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *service) FindUserById(id int64) (*User, error) {

	user, err := u.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *service) FindUserByData(id, name, email string) (*User, error) {

	user, err := u.repository.FindByData(id, name, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *service) DeleteUser(id int64) error {

	user, err := u.repository.FindById(id)
	if err != nil {
		return err
	}

	err = u.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
