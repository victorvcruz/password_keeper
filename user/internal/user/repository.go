package user

import (
	"gorm.io/gorm"
)

type UserRepositoryClient interface {
	Create(user *User) error
	Update(user *User) error
	ExistEmail(email string) bool
	FindById(id int64) (*User, error)
	FindByData(id, name, email string) (*User, error)
	Delete(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(_db *gorm.DB) UserRepositoryClient {
	return &userRepository{
		db: _db,
	}
}

func (u *userRepository) Create(user *User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *User) error {
	err := u.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ExistEmail(email string) bool {
	var exists bool
	u.db.Model(&User{}).
		Select("count(*) > 0").
		Where("email = ? AND deleted_at is null", email).
		Find(&exists)

	return exists
}

func (u *userRepository) FindById(id int64) (*User, error) {
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

func (u *userRepository) Delete(user *User) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindByData(id, name, email string) (*User, error) {
	var user User
	err := u.db.Where("id LIKE ? AND email LIKE ? AND name LIKE ? AND deleted_at is null",
		"%"+id+"%", "%"+email+"%", "%"+name+"%").
		Limit(1).
		Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, nil
	}

	return &user, nil
}
