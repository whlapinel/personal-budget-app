package main

import (
	"time"
)


type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userID"`
	AccountID  int       `json:"accountID"` // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
	Date       time.Time `json:"date"`
	Payee      string    `json:"payee"`
	Amount     float64   `json:"amount"`
	Memo       string    `json:"memo"`
	CategoryID int       `json:"categoryID"` // Simplified version, consider interface{} or struct{} for complex scenarios
}
