package user

import (
	"user.com/internal/platform/database"
)

type RepositoryClient interface {
	Create(user *User) error
	Update(user *User) error
	ExistEmail(email string) bool
	FindById(id int64) (*User, error)
	FindByData(id, name, email string) (*User, error)
	Delete(user *User) error
}

type repository struct {
	db database.Client
}

func NewUserRepository(_db database.Client) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (u *repository) Create(user *User) error {
	err := u.db.DB().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *repository) Update(user *User) error {
	err := u.db.DB().Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *repository) ExistEmail(email string) bool {
	var exists bool
	u.db.DB().Model(&User{}).
		Select("count(*) > 0").
		Where("email = ? AND deleted_at is null", email).
		Find(&exists)

	return exists
}

func (u *repository) FindById(id int64) (*User, error) {
	var user User
	err := u.db.DB().Where("id = ? AND deleted_at is null", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, nil
	}

	return &user, nil
}

func (u *repository) Delete(user *User) error {
	err := u.db.DB().Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *repository) FindByData(id, name, email string) (*User, error) {
	var user User
	err := u.db.DB().Where("id LIKE ? AND email LIKE ? AND name LIKE ? AND deleted_at is null",
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
