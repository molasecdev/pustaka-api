package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func PublisherRoutes(r *gin.RouterGroup) {
	publishers := r.Group("/publishers")

	publishers.POST("/", controllers.CreatePublisher)
	publishers.GET("/", controllers.GetAllPublishers)
	publishers.GET("/:id", controllers.GetPublisherById)
	publishers.PUT("/:id", controllers.UpdatePublisher)
	publishers.DELETE("/:id", controllers.DeletePublisher)
}
