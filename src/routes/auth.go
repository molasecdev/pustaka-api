package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	auth.GET("/validate", controllers.Validate)
}
