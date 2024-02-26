package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreateLanguage(c *gin.Context) {
	db := config.InitConfig()
	var findLanguage models.Language
	var body types.InputLanguage

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	db.Where("LOWER(language) = LOWER(?)", body.Language).First(&findLanguage)
	if findLanguage.Language == body.Language {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language already exist"})
		return
	}

	language := models.Language{
		Language: body.Language,
	}

	errCreate := db.Create(&language)
	if errCreate.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Language created successfully"})
}

func GetAllLanguages(c *gin.Context) {
	db := config.InitConfig()
	var languages []models.Language

	errFLanguages := db.Find(&languages)
	if errFLanguages.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Languages not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": languages})
}

func GetLanguageById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var language models.Language

	errFLanguage := db.Where("id = ?", paramId).First(&language)
	if errFLanguage.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": language})
}

func UpdateLanguage(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body models.Language
	var language models.Language

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFLanguage := db.Where("id = ?", paramId).First(&language)
	if errFLanguage.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found!"})
		return
	}

	if language != (models.Language{}) {
		db.Model(&language).Where("id = ?", paramId).Updates(body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": language})
}

func DeleteLanguage(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var language models.Language

	errFLanguage := db.Where("id = ?", paramId).First(&language)
	if errFLanguage.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found!"})
	}

	db.Where("id = ?", paramId).Delete(&language)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete language successfully!"})
}
