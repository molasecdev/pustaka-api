package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func LanguageRoutes(r *gin.RouterGroup) {
	languages := r.Group("/languages")

	languages.POST("/", controllers.CreateLanguage)
	languages.GET("/", controllers.GetAllLanguages)
	languages.GET("/:id", controllers.GetLanguageById)
	languages.PUT("/:id", controllers.UpdateLanguage)
	languages.DELETE("/:id", controllers.DeleteLanguage)
}
