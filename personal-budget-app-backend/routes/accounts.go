package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"

	"github.com/gin-gonic/gin"
)

func RegisterAccountsRoutes(router *gin.Engine) error {
	router.POST("/accounts", PostAccount)
	router.GET("/accounts/:email", GetAccountsByEmail)
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

func GetAccountsByEmail(c *gin.Context) {
	var account models.Account
	// get accounts
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting accounts"})
		return
	}
	var accounts []models.Account
	for rows.Next() {
		err := rows.Scan(&account.ID, &account.Email, &account.Name, &account.Type, &account.BankName, &account.StartingBalance, &account.Balance)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting accounts"})
			return
		} else {
			accounts = append(accounts, account)
		}
	}
	c.JSON(http.StatusOK, accounts)
}
