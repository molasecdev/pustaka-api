package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func CreatePublisher(c *gin.Context) {
	db := config.InitConfig()
	var findPublisher models.Publisher
	var body types.InputPublisher

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	db.Where("LOWER(publisher) = LOWER(?)", body.Publisher).First(&findPublisher)
	if findPublisher.Publisher == body.Publisher {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Publisher already exist"})
		return
	}

	publisher := models.Publisher{
		Publisher: body.Publisher,
	}

	errCreatePublisher := db.Create(&publisher)
	if errCreatePublisher.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreatePublisher.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Publisher created successfully"})
}

func GetAllPublishers(c *gin.Context) {
	db := config.InitConfig()
	var publishers []models.Publisher

	errFPublisher := db.Find(&publishers)
	if errFPublisher.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": publishers})
}

func GetPublisherById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var publisher models.Publisher

	errFPublisher := db.Where("id = ?", paramId).First(&publisher)
	if errFPublisher.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": publisher})
}

func UpdatePublisher(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body models.Publisher
	var publisher models.Publisher

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFPublisher := db.Where("id = ?", paramId).First(&publisher)
	if errFPublisher.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found!"})
		return
	}

	if publisher != (models.Publisher{}) {
		db.Model(&publisher).Where("id = ?", publisher.ID).Updates(body)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": publisher})
}

func DeletePublisher(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var publisher models.Publisher

	errFPublisher := db.Where("id = ?", paramId).First(&publisher)
	if errFPublisher.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found!"})
	}

	db.Where("id = ?", paramId).Delete(&publisher)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete publisher successfully!"})
}
