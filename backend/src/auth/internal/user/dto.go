package user

import (
	"time"
)

type UserDTO struct {
	Id             int64     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email,omitempty"`
	MasterPassword string    `json:"masterPassword,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}
