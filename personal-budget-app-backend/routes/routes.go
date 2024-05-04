package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	RegisterMonthlyBudgetsRoutes(router)
	RegisterUserRoutes(router)
	RegisterCategoriesRoutes(router)
	RegisterGoalsRoutes(router)
	RegisterAccountsRoutes(router)
	RegisterTransactionsRoutes(router)
	RegisterBudgetPageDataRoutes(router)
	RegisterIncomeAndExpensesRoute(router)
}
