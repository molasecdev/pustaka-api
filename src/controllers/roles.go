package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	db := config.InitConfig()
	var roleExist models.Role
	var rolePayload models.Role

	if err := c.ShouldBindJSON(&rolePayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("LOWER(role) = LOWER(?)", rolePayload.Role).First(&roleExist)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Role has already!"})
		return
	}

	role := models.Role{Role: rolePayload.Role}
	createRole := db.Create(&role)
	if createRole.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createRole.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Role created successfully"})
}

func GetAllRoles(c *gin.Context) {
	db := config.InitConfig()
	var roles []models.Role

	result := db.Find(&roles)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": roles})
}

func GetRoleById(c *gin.Context) {
	db := config.InitConfig()
	var role models.Role
	paramId := c.Param("id")

	result := db.Where("id = ?", paramId).First(&role)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": role})
}

func UpdateRole(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var role models.Role
	var roleInput models.Role

	result := db.Where("id = ?", paramId).First(&role)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	db.Model(&role).Updates(roleInput)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": role})
}

func DeleteRole(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var role models.Role

	result := db.Where("id = ?", paramId).First(&role)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	db.Where("id = ?", paramId).Delete(&role)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete role successfully"})
}
