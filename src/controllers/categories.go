package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	db := config.InitConfig()
	var findCategory models.Category
	var body types.InputCategory

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": errBind.Error()})
		return
	}

	db.Where("LOWER(category) = LOWER(?)", body.Category).First(&findCategory)
	if findCategory.Category == body.Category {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category already exists"})
		return
	}

	category := models.Category{
		Category: body.Category,
	}

	errCreateCategory := db.Create(&category)
	if errCreateCategory.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreateCategory.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Category created successfully"})
}

func GetAllCategories(c *gin.Context) {
	db := config.InitConfig()
	var categories []models.Category

	errFCategory := db.Find(&categories)
	if errFCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
}

func GetCategoryById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var category models.Category

	errFCategory := db.Where("id = ?", paramId).First(&category)
	if errFCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": category})
}

func UpdateCategory(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var category models.Category
	var body models.Category

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFCategory := db.Where("id = ?", paramId).First(&category)
	if errFCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found!"})
		return
	}

	if category != (models.Category{}) {
		db.Model(&category).Where("id = ?", paramId).Updates(body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": category})
}

func DeleteCategory(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var category models.Category

	errFCategory := db.Where("id = ?", paramId).First(&category)
	if errFCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found!"})
		return
	}

	db.Where("id = ?", paramId).Delete(&category)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete category successfully!"})
}
