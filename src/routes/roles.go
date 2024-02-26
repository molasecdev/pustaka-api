package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.RouterGroup) {
	roles := r.Group("/roles")

	roles.POST("/", controllers.CreateRole)
	roles.GET("/", controllers.GetAllRoles)
	roles.GET("/:id", controllers.GetRoleById)
	roles.PUT("/:id", controllers.UpdateRole)
	roles.DELETE("/:id", controllers.DeleteRole)
}
