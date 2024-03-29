package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"personal-budget-app-backend/models"
)

func RegisterMonthlyBudgetRoutes(router *gin.Engine) error {
	router.GET("/monthly-budget/:email/:month/:year", GetMonthlyBudgetHandler)
	return nil
}


func GetMonthlyBudgetHandler(c *gin.Context) {
	// get monthly budget
	email := c.Param("email")
	// coerce month param string into int
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting month from params"})
		return
	}
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting year from params"})
		return
	}
	fmt.Println("email: ", email)
	fmt.Println("month: ", month)
	fmt.Println("year: ", year)
	monthlyBudget, err := models.GetMonthlyBudget(email, month, year)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting monthly budget"})
		return
	}
	c.JSON(http.StatusOK, monthlyBudget)
}
