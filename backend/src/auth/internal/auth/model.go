package auth

import (
	"fmt"
	"time"
)

const schema = "auth_service"

type Auth struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    int64
	Token     string
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (a *Auth) TableName() string {
	return fmt.Sprintf("%s.%s", schema, "auth")
}
