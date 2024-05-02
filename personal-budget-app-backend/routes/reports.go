package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/models/viewmodels"

	"github.com/gin-gonic/gin"
)

func RegisterIncomeAndExpensesRoute(router *gin.Engine) error {
	router.GET("/income-and-expenses/:email", GetIncomeAndExpenses)
	return nil
}

func GetIncomeAndExpenses(c *gin.Context) {
	email := c.Param("email")
	fmt.Println(email)
	data, err := viewmodels.GetIncomeAndExpenses(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
