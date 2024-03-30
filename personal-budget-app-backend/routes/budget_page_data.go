package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/models/viewmodels"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterBudgetPageDataRoutes(router *gin.Engine) error {
	router.GET("/budget-page-data/:email/:month/:year", GetBudgetPageData)
	return nil
}

func GetBudgetPageData(c *gin.Context) {
	email := c.Param("email")
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(email)
	data, err := viewmodels.GetBudgetPageData(email, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
