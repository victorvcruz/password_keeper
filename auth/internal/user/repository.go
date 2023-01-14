package user

import (
	"gorm.io/gorm"
)

type UserRepositoryClient interface {
	UserByEmail(email string) (*User, error)
	UserById(id string) (*User, error)
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
	err := u.db.Where("email = ? AND deleted_at is null", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, nil
	}

	return &user, nil
}

func (u userRepository) UserById(id string) (*User, error) {
	var user User
	err := u.db.Where("id = ? AND deleted_at is null", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, nil
	}

	return &user, nil
}
