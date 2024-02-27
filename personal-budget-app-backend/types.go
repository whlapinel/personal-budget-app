package main

import (
	"time"
)

type User struct {
    ID           int 
    FirstName    string
    LastName     string
    Email        string
}

type BudgetCategory struct {
    ID           int `json:"id"`
	UserID	     int `json:"userID"`
    Name         string `json:"name"`
    Needed       float64 `json:"needed"`
    Assigned     float64 `json:"assigned"`
    Spent        float64 `json:"spent"`
    Available    float64 `json:"available"`
}

type Goal struct {
    ID         string
	UserID	   string
    Name       string
    Amount     float64
    TargetDate time.Time
    Category   BudgetCategory
}

type Transaction struct {
    ID          string
	UserID	    string
    AccountID   string // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
    Date        time.Time
    Payee       string
    Amount      float64
    Memo        string
    CategoryID  string // Simplified version, consider interface{} or struct{} for complex scenarios
}

type Account struct {
    ID       string
	UserID	 string
    Name     string
    Type     AccountType
    BankName string
    Balance  float64
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