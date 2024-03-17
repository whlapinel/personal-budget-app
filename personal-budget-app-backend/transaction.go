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
	Memo       string    `json:"memo"`
	CategoryID int       `json:"categoryID"` // Simplified version, consider interface{} or struct{} for complex scenarios
}

func (t *Transaction) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO transactions (account_id, date, payee, amount, memo, category_id) VALUES ($1, $2, $3, $4, $5, $6)", t.AccountID, t.Date, t.Payee, t.Amount, t.Memo, t.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
