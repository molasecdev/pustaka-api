package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.RouterGroup) {
	AuthRoutes(router)
	UserRoutes(router)
	RoleRoutes(router)
	AuthorRoutes(router)
	PublisherRoutes(router)
	BookselfRoutes(router)
	CategoryRoutes(router)
	LanguageRoutes(router)
	BookRoutes(router)
	LoanRoutes(router)
	FileRoutes(router)
	DashboardRoutes(router)
}
