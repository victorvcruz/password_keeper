package auth

import (
	"auth.com/internal/platform/database"
)

type RepositoryClient interface {
	CreateUserAuth(auth *Auth) (*Auth, error)
	FindByToken(token string) (*Auth, error)
	FindAuthByUser(user int64) (*Auth, error)
	UpdateToken(auth *Auth) error
}
type repository struct {
	db database.Client
}

func NewAuthRepository(_db database.Client) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (a *repository) CreateUserAuth(auth *Auth) (*Auth, error) {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Create(auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *repository) UpdateToken(auth *Auth) error {
	db, err := a.db.Begin()
	defer func() { a.db.CommitOrRollback(db, err) }()

	err = db.Model(auth).Where(&Auth{Id: auth.Id}).Updates(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *repository) FindByToken(token string) (*Auth, error) {
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

func (a *repository) FindAuthByUser(user int64) (*Auth, error) {
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
