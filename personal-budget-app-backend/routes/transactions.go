package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionsRoutes(router *gin.Engine) error {
	router.GET("/transactions/:email", GetTransactionsByEmail)
	router.POST("/transactions", PostTransaction)
	return nil
}

func GetTransactionsByEmail(c *gin.Context) {
	fmt.Println("getTransactionsByEmail")
	var transaction models.Transaction
	// get transactions
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query(`
	SELECT transactions.*, categories.name 
	FROM transactions
	LEFT JOIN categories ON categories.id = transactions.category_id 
	WHERE transactions.email = ?`, email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting transactions"})
		return
	}
	var transactions []models.Transaction
	for rows.Next() {
		var tempDate []uint8
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &tempDate, &transaction.Payee, &transaction.Amount, &transaction.Memo, &transaction.CategoryID, &transaction.Email, &transaction.CategoryName)
		fmt.Println("transaction: ", transaction)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting transactions"})
			return
		} else {
			transaction.Date, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing transaction date"})
				return
			}
			transactions = append(transactions, transaction)
		}
	}
	c.JSON(http.StatusOK, transactions)
}

func PostTransaction(c *gin.Context) {
	var newTransaction models.Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newTransaction.AccountID", newTransaction.AccountID)
	if err := newTransaction.Save(); err != nil {
		fmt.Println("error in newTransaction.Save(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTransaction)
}
