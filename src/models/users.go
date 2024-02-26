package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"id"`
	Firstname string         `gorm:"column:firstname;not null" json:"firstname"`
	Lastname  string         `gorm:"column:lastname;not null" json:"lastname"`
	Birthday  string         `gorm:"column:birthday;not null" json:"birthday"`
	Address   string         `gorm:"column:address;not null" json:"address"`
	Nik       string         `gorm:"column:nik;not null" json:"nik"`
	Phone     string         `gorm:"column:phone;not null" json:"phone"`
	Role_id   uuid.UUID      `gorm:"type:uuid" json:"role_id"`
	Role      Role           `gorm:"foreignKey:Role_id" json:"role"`
	Auth_id   uuid.UUID      `gorm:"type:uuid" json:"auth_id"`
	Auth      Auth           `gorm:"foreignKey:Auth_id" json:"auth"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
