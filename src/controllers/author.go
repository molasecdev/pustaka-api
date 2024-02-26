package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	db := config.InitConfig()
	var findAuthor models.Author
	var body types.InputAuthor

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Where("LOWER(author) = LOWER(?)", body.Author).First(&findAuthor)
	if body.Author == findAuthor.Author {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author already exists!"})
		return
	}

	author := models.Author{
		Author: body.Author,
	}

	errcCreateAuthor := db.Create(&author)
	if errcCreateAuthor.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errcCreateAuthor.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Author created successfully"})
}

func GetAllAuthor(c *gin.Context) {
	db := config.InitConfig()
	var authors []models.Author

	errFAuthor := db.Find(&authors)
	if errFAuthor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": authors})
}

func GetAuthorById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var author models.Author

	errFAuthor := db.Where("id = ?", paramId).First(&author)
	if errFAuthor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": author})
}

func UpdateAuthor(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body models.Author
	var author models.Author

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFAuthor := db.Where("id = ?", paramId).First(&author)
	if errFAuthor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found!"})
		return
	}

	if author != (models.Author{}) {
		db.Model(&author).Where("id = ?", author.ID).Updates(body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": author})
}

func DeleteAuthor(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var author models.Author

	errFAuthor := db.Where("id = ?", paramId).First(&author)
	if errFAuthor.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found!"})
		return
	}

	db.Where("id = ?", paramId).Delete(&author)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete author successfully"})
}
