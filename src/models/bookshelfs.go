package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bookshelfs struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Bookshelf string         `gorm:"column:bookshelf;not null" json:"bookshelf"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (bookshelf *Bookshelfs) BeforeCreate(tx *gorm.DB) (err error) {
	bookshelf.ID = uuid.New()
	bookshelf.CreatedAt = time.Now()
	bookshelf.UpdatedAt = time.Now()
	return
}

func (bookshelf *Bookshelfs) BeforeUpdate(tx *gorm.DB) (err error) {
	bookshelf.UpdatedAt = time.Now()
	return
}
