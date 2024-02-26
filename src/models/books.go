package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID           uuid.UUID      `gorm:"type:uuid" json:"id"`
	Title        string         `gorm:"column:title;not null" json:"title"`
	Description  string         `gorm:"column:description;not null" json:"description"`
	Stock        int            `gorm:"column:stock;not null" json:"stock"`
	Isbn         string         `gorm:"column:isbn;not null" json:"isbn"`
	Year         string         `gorm:"column:year;not null" json:"year"`
	Pages        int            `gorm:"column:pages;not null" json:"pages"`
	Image        string         `gorm:"column:image" json:"image"`
	Author_id    uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
	Author       Author         `gorm:"foreignKey:Author_id" json:"author"`
	Publisher_id uuid.UUID      `gorm:"type:uuid;not null" json:"publisher_id"`
	Publisher    Publisher      `gorm:"foreignKey:Publisher_id" json:"publisher"`
	Category_id  uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Category     Category       `gorm:"foreignKey:Category_id" json:"category"`
	Bookshelf_id uuid.UUID      `gorm:"type:uuid;not null" json:"bookshelf_id"`
	Bookshelf    Bookshelfs     `gorm:"foreignKey:Bookshelf_id" json:"bookshelf"`
	Language_id  uuid.UUID      `gorm:"type:uuid;not null" json:"language_id"`
	Language     Language       `gorm:"foreignKey:Language_id" json:"language"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	return
}

func (book *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	book.UpdatedAt = time.Now()
	return
}
