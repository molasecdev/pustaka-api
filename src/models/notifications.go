package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	Message   string     `gorm:"column:message" json:"message"`
	Read      bool       `gorm:"column:read" json:"read"`
	User_id   uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User      User       `gorm:"foreignKey:user_id" json:"user"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (notification *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	notification.ID = uuid.New()
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()
	return
}

func (notification *Notification) BeforeUpdate(tx *gorm.DB) (err error) {
	notification.UpdatedAt = time.Now()
	return
}
