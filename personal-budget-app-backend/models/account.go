package models

import (
	"personal-budget-app-backend/database"
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

func (a *Account) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO accounts (email, name, type, bank_name, starting_balance, balance) VALUES (?, ?, ?, ?, ?, ?)", a.Email, a.Name, a.Type, a.BankName, a.StartingBalance, a.StartingBalance)
	if err != nil {
		return err
	}
	return nil
}


