package folder

import (
	"gorm.io/gorm"
	"time"
)

type Folder struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*Folder) TableName() string {
	return "folders"
}

func (f *Folder) BeforeCreate(_ *gorm.DB) (err error) {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return nil
}

func (f *Folder) BeforeUpdate(_ *gorm.DB) (err error) {
	f.UpdatedAt = time.Now()
	return nil
}

func (f *Folder) ToResponse() Response {
	return Response{
		ID:        f.ID,
		UserID:    f.UserID,
		Name:      f.Name,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}
