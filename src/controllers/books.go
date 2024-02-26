package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	db := config.InitConfig()
	var body types.InputBook
	var author models.Author
	var publisher models.Publisher
	var category models.Category
	var bookshelf models.Bookshelfs
	var language models.Language

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFAuthor := db.Where("LOWER(author) = LOWER(?)", body.Author).First(&author)
	errFPublisher := db.Where("LOWER(publisher) = LOWER(?)", body.Publisher).First(&publisher)
	errFCategory := db.Where("LOWER(category) = LOWER(?)", body.Category).First(&category)
	errFBookshelf := db.Where("LOWER(bookshelf) = LOWER(?)", body.Bookshelf).First(&bookshelf)
	errFLanguage := db.Where("LOWER(language) = LOWER(?)", body.Language).First(&language)

	if errFAuthor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	if errFPublisher.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}
	if errFCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if errFBookshelf.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookshelf not found"})
		return
	}
	if errFLanguage.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found"})
		return
	}

	book := models.Book{
		Title:        body.Title,
		Description:  body.Description,
		Stock:        body.Stock,
		Isbn:         body.Isbn,
		Year:         body.Year,
		Pages:        body.Pages,
		Image:        body.Image,
		Author_id:    author.ID,
		Publisher_id: publisher.ID,
		Category_id:  category.ID,
		Bookshelf_id: bookshelf.ID,
		Language_id:  language.ID,
	}

	createBook := db.Create(&book)
	if createBook.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book created successfully"})
}

func GetAllBooks(c *gin.Context) {
	db := config.InitConfig()
	var books []models.Book

	errFBooks := db.Preload("Author").Preload("Publisher").Preload("Category").Preload("Bookshelf").Preload("Language").Find(&books)
	if errFBooks.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Books not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": books})
}

func GetBookById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var book models.Book

	errFBook := db.Where("id = ?", paramId).Preload("Author").Preload("Publisher").Preload("Category").Preload("Bookshelf").Preload("Language").First(&book)
	if errFBook.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": book})
}

func UpdateBook(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body types.UpdateBook
	var book models.Book

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	if errFBook := db.Where("id = ?", paramId).First(&book).Error; errFBook != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book != (models.Book{}) {
		db.Model(&book).Where("id = ?", paramId).Updates(body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": book})
}

func DeleteBook(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var book models.Book

	if errFBook := db.Where("id = ?", paramId).First(&book).Error; errFBook != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Where("id = ?", paramId).Delete(&book)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete book successfully!"})
}
