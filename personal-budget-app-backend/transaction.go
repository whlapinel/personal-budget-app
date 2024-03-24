package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID           int       `json:"id"`
	AccountID    int       `json:"accountID"` // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
	Date         time.Time `json:"date"`
	Payee        string    `json:"payee"`
	Amount       int       `json:"amount"`     // in cents not dollars
	Memo         *string   `json:"memo"`       // pointer so value can be nil
	CategoryID   *int      `json:"categoryID"` // pointer so value can be nil
	CategoryName *string   `json:"categoryName"`
}

func (t *Transaction) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO transactions (account_id, date, payee, amount, memo, category_id) VALUES (?, ?, ?, ?, ?, ?)", t.AccountID, t.Date, t.Payee, t.Amount, t.Memo, t.CategoryID)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_account_balance(?)", t.AccountID)
	if err != nil {
		return err
	}
	fmt.Println("Transaction created and account balance updated for account: ", t.AccountID)
	return nil
}

func getTransactionsByAccountID(c *gin.Context) {
	fmt.Println("getTransactionsByAccountID")
	var transaction Transaction
	// get transactions
	accountID := c.Param("accountID")
	fmt.Println("accountID: ", accountID)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query(`
	SELECT * 
	FROM transactions 
	WHERE transactions.account_id = ?`, accountID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting transactions"})
		return
	}
	var transactions []Transaction
	for rows.Next() {
		var tempDate []uint8
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &tempDate, &transaction.Payee, &transaction.Amount, &transaction.Memo, &transaction.CategoryID)
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

func postTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newTransaction.AccountID", newTransaction.AccountID)
	if err := newTransaction.create(); err != nil {
		fmt.Println("error in newTransaction.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTransaction)
}
