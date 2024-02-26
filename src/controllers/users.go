package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db := config.InitConfig()
	var users []models.User

	errFUser := db.Preload("Role").Preload("Auth").Find(&users)
	if errFUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func GetUserById(c *gin.Context) {
	db := config.InitConfig()
	var user models.User
	paramId := c.Param("id")

	errFUser := db.Where("id = ?", paramId).Preload("Role").Preload("Auth").First(&user)
	if errFUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func UpdateUser(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var user models.User
	var auth models.Auth
	var body types.UpdateUser

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFUser := db.Where("id = ?", paramId).First(&user)
	if errFUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	errFAuth := db.Where("id = ?", user.Auth_id).First(&auth)
	if errFAuth.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auth not found"})
		return
	}

	if auth != (models.Auth{}) && (body.Email != "" || body.Password != "") {
		db.Model(&auth).Where("id = ?", auth.ID).Updates(models.Auth{Email: body.Email, Password: body.Password})
	}

	if user != (models.User{}) {
		db.Model(&user).Where("id = ?", user.ID).Updates(body)
	}

	var updatedUser models.User
	db.Where("id = ?", paramId).Preload("Auth").Preload("Role").First(&updatedUser)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": updatedUser})
}

func DeleteUser(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var user models.User
	var auth models.Auth

	errFUser := db.Where("id = ?", paramId).First(&user)
	if errFUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.Where("id = ?", paramId).Delete(&user)
	db.Where("id = ?", user.Auth_id).Delete(&auth)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully"})
}
