package main

import (
	"fmt"
	"net/http"

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

func (a *Account) create() error {
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
	// since I just added a stored procedure to update account balances, I
	// should not need any of the below commented out code.  Keeping just in case.

	// get account balances by retrieving sum of transactions for each account,
	// add to each account struct instance
	// for i, account := range accounts {
	// 	// make sure there are transactions before running query
	// 	rows, err := db.Query("SELECT COUNT(*) FROM transactions WHERE account_id = ?", account.ID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
	// 		return
	// 	}
	// 	var count int
	// 	for rows.Next() {
	// 		err := rows.Scan(&count)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
	// 			return
	// 		}
	// 		if count == 0 {
	// 			accounts[i].Balance = account.StartingBalance
	// 			continue
	// 		} else {
	// 			// get sum of transactions
	// 			rows, err = db.Query("SELECT SUM(amount) FROM transactions WHERE account_id = ?", account.ID)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 				c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
	// 				return
	// 			}
	// 			var balance int
	// 			for rows.Next() {
	// 				err := rows.Scan(&balance)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 					c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting account balances"})
	// 					return
	// 				}
	// 				accounts[i].Balance = balance + account.StartingBalance
	// 			}
	// 		}
	// 	}
	// }
	c.JSON(http.StatusOK, accounts)
}

func postAccount(c *gin.Context) {
	var newAccount Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newAccount)
	if err := newAccount.create(); err != nil {
		fmt.Println("error in newAccount.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAccount)
}
