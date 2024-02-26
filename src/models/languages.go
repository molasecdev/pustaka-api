package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Language  string         `gorm:"column:language;not null" json:"language"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (language *Language) BeforeCreate(tx *gorm.DB) (err error) {
	language.ID = uuid.New()
	language.CreatedAt = time.Now()
	language.UpdatedAt = time.Now()
	return
}

func (language *Language) BeforeUpdate(tx *gorm.DB) (err error) {
	language.UpdatedAt = time.Now()
	return
}
