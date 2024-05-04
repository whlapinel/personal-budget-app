package models

import (
	"log"
	"personal-budget-app-backend/database"
	"time"
)

type Account struct {
	ID              int         `json:"id"`
	Email           string      `json:"email"`
	Name            string      `json:"name"`
	Type            AccountType `json:"type"`
	BankName        string      `json:"bankName"`
	StartingBalance int         `json:"startingBalance"`
	StartingDate    time.Time   `json:"startingDate"`
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
	_, err := db.Exec("INSERT INTO accounts (email, name, type, bank_name, starting_balance, starting_date, balance) VALUES (?, ?, ?, ?, ?, ?, ?)", a.Email, a.Name, a.Type, a.BankName, a.StartingBalance, a.StartingDate, a.StartingBalance)
	if err != nil {
		return err
	}
	return nil
}

func GetAccounts(email string) ([]Account, error) {
	log.Println("GetAccounts email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	var accounts []Account
	for rows.Next() {
		var account Account
		var tempDate []uint8
		err := rows.Scan(&account.ID, &account.Email, &account.Name, &account.Type, &account.BankName, &account.StartingBalance, &tempDate, &account.Balance)
		if err != nil {
			return nil, err
		}
		account.StartingDate, err = time.Parse("2006-01-02", string(tempDate))
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
