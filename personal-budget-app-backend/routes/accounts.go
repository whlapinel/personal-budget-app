package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/models"

	"github.com/gin-gonic/gin"
)

func RegisterAccountsRoutes(router *gin.Engine) error {
	router.POST("/accounts", PostAccount)
	router.GET("/accounts/:email", GetAccountsHandler)
	return nil
}

func PostAccount(c *gin.Context) {
	var newAccount models.Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newAccount)
	if err := newAccount.Save(); err != nil {
		fmt.Println("error in newAccount.Save(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAccount)
}

func GetAccountsHandler(c *gin.Context) {
	// get accounts
	email := c.Param("email")
	fmt.Println("email: ", email)
	accounts, err := models.GetAccounts(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}
