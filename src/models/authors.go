package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Author    string         `gorm:"column:author;not null" json:"author"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (author *Author) BeforeCreate(tx *gorm.DB) (err error) {
	author.ID = uuid.New()
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()
	return
}

func (author *Author) BeforeUpdate(tx *gorm.DB) (err error) {
	author.UpdatedAt = time.Now()
	return
}
