package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.RouterGroup) {
	notifications := r.Group("/notifications")

	notifications.POST("/", controllers.CreateNotification)
	notifications.GET("/", controllers.GetAllNotifications)
	notifications.PUT("/:id", controllers.UpdateNotification)
	notifications.DELETE("/:id", controllers.DeleteNotification)
}
