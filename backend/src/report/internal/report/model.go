package report

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	Id          int `gorm:"primaryKey"`
	UserId      int64
	VaultId     *int64
	Action      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (r *Report) TableName() string {
	return "reports"
}

func (r *Report) BeforeCreate(_ *gorm.DB) (err error) {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Report) BeforeUpdate(_ *gorm.DB) (err error) {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Report) FillFields(action string, userId, vaultId int64, description string) {
	now := time.Now()

	r.Action = action
	r.UserId = userId
	r.Description = description
	r.CreatedAt = now
	r.UpdatedAt = now

	if vaultId == 0 {
		r.VaultId = nil
	} else {
		r.VaultId = &vaultId
	}
}
