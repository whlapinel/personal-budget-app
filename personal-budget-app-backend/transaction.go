package main

import (
	"fmt"
	"net/http"
	"time"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID           int       `json:"id"`
	AccountID    int       `json:"accountID"` // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
	Date         time.Time `json:"date"`
	Payee        string    `json:"payee"`
	Amount       int       `json:"amount"`       // in cents not dollars
	Memo         *string   `json:"memo"`         // pointer so value can be nil
	CategoryID   *int      `json:"categoryID"`   // pointer so value can be nil
	CategoryName *string   `json:"categoryName"` // stored in DB under categories.name
	Email        string    `json:"email"`
}

func createTransactionTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE transactions (
			id int AUTO_INCREMENT PRIMARY KEY,
			account_id int,
			date datetime,
			payee VARCHAR(100),
			amount int,
			memo VARCHAR(100),
			category_id int,
			email VARCHAR(100),
			FOREIGN KEY (email) REFERENCES users(email),
			FOREIGN KEY (account_id) REFERENCES accounts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (t *Transaction) Save() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES (?, ?, ?, ?, ?, ?, ?)", t.AccountID, t.Date, t.Payee, t.Amount, t.Memo, t.CategoryID, t.Email)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_account_balance(?, ?)", t.AccountID, t.Amount)
	if err != nil {
		return err
	}
	fmt.Println("Transaction created and account balance updated for account: ", t.AccountID)
	return nil
}

func getTransactionsByEmail(c *gin.Context) {
	fmt.Println("getTransactionsByEmail")
	var transaction Transaction
	// get transactions
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
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
	var transactions []Transaction
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

func postTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newTransaction.AccountID", newTransaction.AccountID)
	if err := newTransaction.Save(); err != nil {
		fmt.Println("error in newTransaction.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTransaction)
}
