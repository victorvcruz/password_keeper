package folder

import (
	"gorm.io/gorm"
	"time"
)

type Folder struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Folder) TableName() string {
	return "folders"
}
