package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.RouterGroup) {
	books := r.Group("/books")

	books.POST("/", controllers.CreateBook)
	books.GET("/", controllers.GetAllBooks)
	books.GET("/:id", controllers.GetBookById)
	books.PUT("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)
}
