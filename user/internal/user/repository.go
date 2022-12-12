package user

import (
	"gorm.io/gorm"
)

type UserRepositoryClient interface {
	Create(user *User) error
	ExistEmail(email string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(_db *gorm.DB) UserRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u userRepository) Create(user *User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) Delete(user *User) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) ExistEmail(email string) bool {
	var exists bool
	u.db.Model(&User{}).
		Select("count(*) > 0").
		Where("email = ? AND deleted_at is null", email).
		Find(&exists)

	return exists
}
