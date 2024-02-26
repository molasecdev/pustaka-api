package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.RouterGroup) {
	file := r.Group("/files")

	file.POST("/uploads", controllers.UploadFile)
	file.GET("/:filename", controllers.GetFile)
	file.PUT("/:filename", controllers.UpdateFile)
	file.DELETE("/:filename", controllers.DeleteFile)
}
