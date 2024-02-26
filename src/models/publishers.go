package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Publisher struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Publisher string         `gorm:"column:publisher;not null" json:"publisher"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (publisher *Publisher) BeforeCreate(tx *gorm.DB) (err error) {
	publisher.ID = uuid.New()
	publisher.CreatedAt = time.Now()
	publisher.UpdatedAt = time.Now()
	return
}

func (publisher *Publisher) BeforeUpdate(tx *gorm.DB) (err error) {
	publisher.UpdatedAt = time.Now()
	return
}
