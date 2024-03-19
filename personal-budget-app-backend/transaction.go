package main

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"accountID"` // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
	Date       time.Time `json:"date"`
	Payee      string    `json:"payee"`
	Amount     float64   `json:"amount"`
	Memo       *string    `json:"memo"`    // pointer so value can be nil
	CategoryID *int      `json:"categoryID"` // pointer so value can be nil
}

func (t *Transaction) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO transactions (account_id, date, payee, amount, memo, category_id) VALUES (?, ?, ?, ?, ?, ?)", t.AccountID, t.Date, t.Payee, t.Amount, t.Memo, t.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
