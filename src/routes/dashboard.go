package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.RouterGroup) {
	dashboard := r.Group("/dashboard")

	dashboard.GET("/get-book-count", controllers.GetBooksCount)
	dashboard.GET("/get-active-loan", controllers.GetActiveLoanSummary)
	dashboard.GET("/get-user-statistic", controllers.GetUserStatistics)
}
