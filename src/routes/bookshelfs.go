package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func BookselfRoutes(r *gin.RouterGroup) {
	bookselfs := r.Group("/bookshelfs")

	bookselfs.POST("/", controllers.CreateBookshelf)
	bookselfs.GET("/", controllers.GetAllBookshelfs)
	bookselfs.GET("/:id", controllers.GetBookshelfById)
	bookselfs.PUT("/:id", controllers.UpdateBookshelf)
	bookselfs.DELETE("/:id", controllers.DeleteBookshelf)
}
