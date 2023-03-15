package service

import (
	"gorm.io/gorm"
	"time"
)

type Internal struct {
	Id        int64 `gorm:"primaryKey"`
	Service   string
	Token     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *Internal) FillFields(service, token, password string) {
	now := time.Now()

	u.Service = service
	u.Password = password
	u.Token = token
	u.CreatedAt = now
	u.UpdatedAt = now
}
