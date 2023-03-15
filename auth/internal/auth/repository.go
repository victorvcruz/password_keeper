package auth

import "gorm.io/gorm"

type AuthRepositoryClient interface {
	Create(user *Auth) error
	FindByToken(token string) (*Auth, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewAuthRepository(_db *gorm.DB) AuthRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u *userRepository) Create(auth *Auth) error {
	err := u.db.Create(auth).Error
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
