package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreateBookshelf(c *gin.Context) {
	db := config.InitConfig()
	var findBookshelf models.Bookshelfs
	var body types.InputBookshelf

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	db.Where("LOWER(bookshelf) = LOWER(?)", body.Bookshelf).First(&findBookshelf)
	if findBookshelf.Bookshelf == body.Bookshelf {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookshelf already exists"})
		return
	}

	bookshelf := models.Bookshelfs{
		Bookshelf: body.Bookshelf,
	}

	errCreateBookshelf := db.Create(&bookshelf)
	if errCreateBookshelf.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreateBookshelf.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Bookshelf added successfully"})
}

func GetAllBookshelfs(c *gin.Context) {
	db := config.InitConfig()
	var bookshelfs []models.Bookshelfs

	errFBookshelf := db.Find(&bookshelfs)
	if errFBookshelf.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookshelf not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": bookshelfs})
}

func GetBookshelfById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var bookshelf models.Bookshelfs

	errFBookshelf := db.Where("id = ?", paramId).First(&bookshelf)
	if errFBookshelf.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookshelf not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": bookshelf})
}

func UpdateBookshelf(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body models.Bookshelfs
	var bookshelf models.Bookshelfs

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFBookshelf := db.Where("id = ?", paramId).First(&bookshelf)
	if errFBookshelf.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookshelf not found!"})
		return
	}

	if bookshelf != (models.Bookshelfs{}) {
		db.Model(&bookshelf).Where("id = ?", bookshelf.ID).Updates(&body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": bookshelf})
}

func DeleteBookshelf(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var bookshelf models.Bookshelfs

	errFBookshelf := db.Where("id = ?", paramId).First(&bookshelf)
	if errFBookshelf.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bookshelf not found!"})
	}

	db.Where("id = ?", paramId).Delete(&bookshelf)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete bookshelf successfully!"})
}
