package main

import (
)

type Account struct {
	ID       int         `json:"id"`
	UserID   int         `json:"userID"`
	Name     string      `json:"name"`
	Type     AccountType `json:"type"`
	BankName string      `json:"bankName"`
	Balance  float64     `json:"balance"`
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
