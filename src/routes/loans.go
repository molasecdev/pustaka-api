package routes

import (
	"pustaka-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func LoanRoutes(r *gin.RouterGroup) {
	loans := r.Group("/loans")

	loans.POST("/", controllers.CreateLoan)
	loans.GET("/", controllers.GetAllLoans)
	loans.POST("/exports", controllers.ExportLoans)
	loans.GET("/:id", controllers.GetLoanById)
	loans.PUT("/:id", controllers.UpdateLoan)
	loans.DELETE("/:id", controllers.DeleteLoan)
}
