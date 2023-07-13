package auth

import (
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    int64
	Token     string
	CreatedAt time.Time
	ExpiredAt time.Time
}

type AuthApi struct {
	Id          int64 `gorm:"primaryKey"`
	Service     int64
	ServiceConn int64
	ApiToken    string
	Token       string
	CreatedAt   time.Time
	ExpiredAt   time.Time
}

type Service struct {
	Id        int64 `gorm:"primaryKey"`
	Name      string
	ApiToken  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
