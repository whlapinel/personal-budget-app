package main

import ()

type Account struct {
	ID              int         `json:"id"`
	Email           string      `json:"email"`
	Name            string      `json:"name"`
	Type            AccountType `json:"type"`
	BankName        string      `json:"bankName"`
	StartingBalance float64     `json:"startingBalance"`
	Balance         float64     `json:"balance"`
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
	_, err := db.Exec("INSERT INTO accounts (email, name, type, bank_name, starting_balance) VALUES (?, ?, ?, ?, ?)", a.Email, a.Name, a.Type, a.BankName, a.StartingBalance)
	if err != nil {
		return err
	}
	return nil
}
