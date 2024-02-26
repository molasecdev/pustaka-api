package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.RouterGroup) {
	authors := r.Group("/authors")

	authors.POST("/", controllers.CreateAuthor)
	authors.GET("/", controllers.GetAllAuthor)
	authors.GET("/:id", controllers.GetAuthorById)
	authors.PUT("/:id", controllers.UpdateAuthor)
	authors.DELETE("/:id", controllers.DeleteAuthor)
}
