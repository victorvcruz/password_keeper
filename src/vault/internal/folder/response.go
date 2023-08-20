package folder

import (
	"time"
)

type Response struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
