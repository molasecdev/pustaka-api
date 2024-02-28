package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func SaveNotification(notification models.Notification) error {
	db := config.InitConfig()

	notification.Read = false

	if err := db.Create(&notification).Error; err != nil {
		return err
	}

	return nil
}

func CreateNotification(c *gin.Context) {
	var body types.InputNotification

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	notif := models.Notification{
		User_id: body.User_id,
		Message: body.Message,
	}

	if err := SaveNotification(notif); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create notification successfully!"})
}

func GetAllNotifications(c *gin.Context) {
	db := config.InitConfig()
	var notifications []models.Notification

	if errFNotif := db.Preload("User").Find(&notifications).Error; errFNotif != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notifications not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": notifications})
}

func UpdateNotification(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body types.UpdateNotification
	var notification models.Notification

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	if errFNotif := db.Where("id = ?", paramId).Preload("User").First(&notification).Error; errFNotif != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found!"})
		return
	}

	if notification != (models.Notification{}) {
		expired := notification.UpdatedAt.AddDate(0, 0, 1) // akan dihapus 1 harinya setelah sudah dibaca
		var dataNotif interface{}

		if !notification.Read && notification.DeletedAt == nil {
			dataNotif = models.Notification{
				Read:      body.Read,
				DeletedAt: &expired,
			}
		}

		db.Model(&notification).Where("id = ?", paramId).Updates(dataNotif)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": notification})
}

func DeleteNotification(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var notification models.Notification

	if errFNotif := db.Where("id = ?", paramId).First(&notification).Error; errFNotif != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found!"})
		return
	}

	db.Where("id = ?", paramId).Delete(&notification)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete notification successfully!"})
}
