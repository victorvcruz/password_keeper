package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id             int64 `gorm:"primaryKey"`
	Name           string
	Email          string
	MasterPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) FillFields(name, email, masterPassword string) {
	u.Name = name
	u.Email = email
	u.MasterPassword = masterPassword
}

func (u *User) ToUpdate(name, email, masterPassword string) {
	now := time.Now()

	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if masterPassword != "" {
		u.MasterPassword = masterPassword
	}

	u.UpdatedAt = now
}
