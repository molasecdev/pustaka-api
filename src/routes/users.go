package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")

	users.GET("/", controllers.GetAllUsers)
	users.GET("/:id", controllers.GetUserById)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
