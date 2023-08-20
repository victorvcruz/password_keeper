package vault

import (
	"gorm.io/gorm"
	"time"
)

type Vault struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64
	FolderID  uint64
	Username  string
	Name      string
	Password  string
	URL       string
	Notes     string
	Favorite  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*Vault) TableName() string {
	return "vaults"
}

func (v *Vault) BeforeCreate(_ *gorm.DB) (err error) {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return nil
}

func (v *Vault) BeforeUpdate(_ *gorm.DB) (err error) {
	v.UpdatedAt = time.Now()
	return nil
}

func (v *Vault) ToResponse() Response {
	return Response{
		ID:        v.ID,
		UserID:    v.UserID,
		FolderID:  v.FolderID,
		Username:  v.Username,
		Name:      v.Name,
		Password:  v.Password,
		URL:       v.URL,
		Notes:     v.Notes,
		Favorite:  v.Favorite,
		CreatedAt: &v.CreatedAt,
		UpdatedAt: &v.UpdatedAt,
	}
}
