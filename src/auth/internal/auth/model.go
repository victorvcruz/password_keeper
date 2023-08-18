package auth

import (
	"fmt"
	"gorm.io/gorm"
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

func (s *Auth) TableName() string {
	return fmt.Sprintf("%s.%s", schema, "auth")
}

type AuthService struct {
	Id          int64 `gorm:"primaryKey"`
	Service     int64
	ServiceConn int64
	ApiToken    string
	Token       string
	CreatedAt   time.Time
	ExpiredAt   time.Time
}

func (s *AuthService) TableName() string {
	return fmt.Sprintf("%s.%s", schema, "auth_service")
}

type Service struct {
	Id        int64 `gorm:"primaryKey"`
	Name      string
	ApiToken  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (s *Service) TableName() string {
	return fmt.Sprintf("%s.%s", schema, "service")
}
