package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loan struct {
	ID          uuid.UUID      `gorm:"type:uuid" json:"id"`
	Start_date  time.Time      `gorm:"column:start_date" json:"start_date"`
	End_date    time.Time      `gorm:"column:end_date" json:"end_date"`
	Note        string         `gorm:"column:note" json:"note"`
	Status      string         `gorm:"column:status" json:"status"`
	Return_date *time.Time     `gorm:"column:return_date" json:"return_date"`
	Penalty     int            `gorm:"column:penalty" json:"penalty"`
	User_id     uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	User        User           `gorm:"foreignKey:User_id" json:"user"`
	Book_id     uuid.UUID      `gorm:"type:uuid" json:"book_id"`
	Book        Book           `gorm:"foreignKey:Book_id" json:"book"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (loan *Loan) BeforeCreate(tx *gorm.DB) (err error) {
	loan.ID = uuid.New()
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()
	return
}

func (loan *Loan) BeforeUpdate(tx *gorm.DB) (err error) {
	loan.UpdatedAt = time.Now()
	return
}
