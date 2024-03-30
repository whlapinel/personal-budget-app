package routes

import (
	"net/http"
	"personal-budget-app-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterMonthlyBudgetsRoutes(router *gin.Engine) error {
	router.GET("/monthly-budgets/:email/:month/:year", GetMonthlyBudgetsHandler)
	router.POST("/monthly-budgets", PostMonthlyBudgetsHandler)
	return nil
}

func GetMonthlyBudgetsHandler(c *gin.Context) {
	// get monthly budget
	email := c.Param("email")
	// coerce month param string into int
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	monthlyBudgets, err := models.GetMonthlyBudgets(email, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, monthlyBudgets)
}

func PostMonthlyBudgetsHandler(c *gin.Context) {
	var newMonthlyBudget models.MonthlyBudget
	if err := c.BindJSON(&newMonthlyBudget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := newMonthlyBudget.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMonthlyBudget)
}
