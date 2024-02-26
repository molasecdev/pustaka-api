package controllers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	r := c.Request
	maxSize := int64(5 * 1024 * 1024) // 5 MB

	dir, errDir := os.Getwd()
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// get file
	if err := r.ParseMultipartForm(1024); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	file, fileHeader, errFile := r.FormFile("file")
	if errFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFile.Error()})
		return
	}
	defer file.Close()

	// type & size validations
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only image files are allowed!"})
		return
	}
	if fileHeader.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large!"})
		return
	}

	// generate UUID for the filename & save file
	fileExt := filepath.Ext(fileHeader.Filename)
	filename := uuid.New().String() + fileExt
	fileLocation := filepath.Join(dir, "files/img", filename)

	targetFile, errCreate := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}
	defer targetFile.Close()

	if _, errCopy := io.Copy(targetFile, file); errCopy != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCopy.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": gin.H{"filename": filename}})
}

func GetFile(c *gin.Context) {
	filename := c.Param("filename")

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileLocation := filepath.Join(dir, "files/img", filename)

	// check file
	_, err = os.Stat(fileLocation)
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found!"})
		return
	}

	c.File(fileLocation)
}

func UpdateFile(c *gin.Context) {
	filename := c.Param("filename")
	maxSize := int64(5 * 1024 * 1024) // 5 MB

	dir, errDir := os.Getwd() // get location
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// delete old file
	oldFileLocation := filepath.Join(dir, "files/img", filename)
	if errFileLocation := os.Remove(oldFileLocation); errFileLocation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFileLocation.Error()})
		return
	}

	// get newFile
	newFile, newFileHeader, errFile := c.Request.FormFile("file")
	if errFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFile.Error()})
		return
	}
	defer newFile.Close()

	// type & size validations
	contentType := newFileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only image files are allowed!"})
		return
	}
	if newFileHeader.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large!"})
		return
	}

	// Generate UUID for the filename & save new file
	fileExt := filepath.Ext(newFileHeader.Filename)
	newFilename := uuid.New().String() + fileExt
	newFileLocation := filepath.Join(dir, "files/img", newFilename)

	newTargetFile, errCreate := os.OpenFile(newFileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}
	defer newTargetFile.Close()

	if _, errCopy := io.Copy(newTargetFile, newFile); errCopy != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCopy.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "filename": newFilename})
}

func DeleteFile(c *gin.Context) {
	filename := c.Param("filename")

	dir, errDir := os.Getwd() // get location
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// delete old file
	fileLocation := filepath.Join(dir, "files/img", filename)
	if errFileLocation := os.Remove(fileLocation); errFileLocation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFileLocation.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "File deleted successfully"})
}
