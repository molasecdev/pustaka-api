package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	categories := r.Group("/categories")

	categories.POST("/", controllers.CreateCategory)
	categories.GET("/", controllers.GetAllCategories)
	categories.GET("/:id", controllers.GetCategoryById)
	categories.PUT("/:id", controllers.UpdateCategory)
	categories.DELETE("/:id", controllers.DeleteCategory)
}
