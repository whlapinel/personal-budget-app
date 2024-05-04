package routes

import (
	"net/http"
	"personal-budget-app-backend/models/viewmodels"

	"github.com/gin-gonic/gin"
)

func RegisterNetWorthReportRoutes(router *gin.Engine) error {
	router.GET("/net-worth-report/:email", GetNetWorthReportHandler)
	return nil
}

func GetNetWorthReportHandler(c *gin.Context) {
	// get net worth report
	email := c.Param("email")
	netWorthReport, err := viewmodels.GetNetWorthReport(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, netWorthReport)
}
