package auth

import (
	"time"
)

type Auth struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    int64
	Token     string
	CreatedAt time.Time
	ExpiredAt time.Time
}
