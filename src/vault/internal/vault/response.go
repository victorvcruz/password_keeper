package vault

import (
	"time"
)

type Response struct {
	ID        uint64     `json:"id"`
	UserID    uint64     `json:"user_id"`
	FolderID  uint64     `json:"folder_id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	URL       string     `json:"url"`
	Notes     string     `json:"notes"`
	Favorite  bool       `json:"favorite"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
