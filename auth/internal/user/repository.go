package user

import (
	"gorm.io/gorm"
)

type UserRepositoryClient interface {
	UserByEmail(email string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(_db *gorm.DB) UserRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u userRepository) UserByEmail(email string) (*User, error) {
	var user User
	err := u.db.Find(&user).Where("email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
