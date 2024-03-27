package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Account struct {
	ID              int         `json:"id"`
	Email           string      `json:"email"`
	Name            string      `json:"name"`
	Type            AccountType `json:"type"`
	BankName        string      `json:"bankName"`
	StartingBalance int         `json:"startingBalance"`
	Balance         int         `json:"balance"`
}

type AccountType string

const (
	Checking   AccountType = "checking"
	Savings    AccountType = "savings"
	Credit     AccountType = "credit"
	Loan       AccountType = "loan"
	Investment AccountType = "investment"
	Other      AccountType = "other"
)

func createAccountTable(db *sql.DB) (sql.Result, error) {

	query :=
		`CREATE TABLE accounts (
			id int AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100),
			name VARCHAR(100),
			type VARCHAR(100),
			bank_name VARCHAR(100),
			starting_balance int,
			balance int,
			FOREIGN KEY (email) REFERENCES users(email)
			);`
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (a *Account) Save() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO accounts (email, name, type, bank_name, starting_balance, balance) VALUES (?, ?, ?, ?, ?, ?)", a.Email, a.Name, a.Type, a.BankName, a.StartingBalance, a.StartingBalance)
	if err != nil {
		return err
	}
	return nil
}

func getAccountsByEmail(c *gin.Context) {
	var account Account
	// get accounts
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting accounts"})
		return
	}
	var accounts []Account
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

func postAccount(c *gin.Context) {
	var newAccount Account
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
