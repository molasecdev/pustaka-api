package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Email     string         `gorm:"column:email;not null" json:"email"`
	Password  string         `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	auth.ID = uuid.New()
	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()
	return
}

func (auth *Auth) BeforeUpdate(tx *gorm.DB) (err error) {
	auth.UpdatedAt = time.Now()
	return
}
