package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionsRoutes(router *gin.Engine) error {
	router.GET("/transactions/:email", GetTransactionsByEmail)
	router.POST("/transactions", PostTransaction)
	router.DELETE("/transactions/:id", DeleteTransaction)
	router.PATCH("/transactions/:id", UpdateTransaction)
	return nil
}

func UpdateTransaction(c *gin.Context) {
	fmt.Println("patchTransaction")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oldTransaction := models.Transaction{ID: id}
	var newTransaction *models.Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = oldTransaction.Update(newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTransaction)
}

func DeleteTransaction(c *gin.Context) {
	fmt.Println("deleteTransaction")
	id := c.Param("id")
	fmt.Println("id: ", id)
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting transaction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}

func GetTransactionsByEmail(c *gin.Context) {
	fmt.Println("getTransactionsByEmail")
	// get transactions
	email := c.Param("email")
	fmt.Println("email: ", email)
	transactions, err := models.GetTransactionsByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
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
