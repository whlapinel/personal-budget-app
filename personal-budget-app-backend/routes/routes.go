package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	RegisterMonthlyBudgetRoutes(router)
	RegisterUserRoutes(router)
	RegisterCategoriesRoutes(router)
	RegisterAssignmentsRoutes(router)
	RegisterGoalsRoutes(router)
	RegisterAccountsRoutes(router)
}