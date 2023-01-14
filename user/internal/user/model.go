package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id             int `gorm:"primaryKey"`
	Name           string
	Email          string
	MasterPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) FillFields(name, email, masterPassword string) {
	now := time.Now()

	u.Name = name
	u.Email = email
	u.MasterPassword = masterPassword
	u.CreatedAt = now
	u.UpdatedAt = now
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
