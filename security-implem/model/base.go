package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string          `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt *time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`
}

// Hook to set UUID before creating a new record
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
