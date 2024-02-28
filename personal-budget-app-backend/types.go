package main

import (
	"time"
)

type User struct {
    ID           int        `json:"id"`
    UserName     string     `json:"userName"`
    Password     string     `json:"password"`
    FirstName    string     `json:"firstName"`
    LastName     string     `json:"lastName"`
    Email        string     `json:"email"`
}

// user methods
func (u *User) CreateAccount(a Account)(Account) {
    newAccount := Account{
        Name: a.Name,
        Type: a.Type,
        BankName: a.BankName,
        Balance: a.Balance,
        UserID: u.ID,
    }    
    return newAccount
}


type BudgetCategory struct {
    ID           int        `json:"id"`
	UserID	     int        `json:"userID"`
    Name         string     `json:"name"`
}

type Goal struct {
    ID         string           `json:"id"`
	UserID	   string           `json:"userID"`
    Name       string           `json:"name"`
    Amount     float64          `json:"amount"`
    TargetDate time.Time        `json:"targetDate"`
    Category   BudgetCategory   `json:"category"`
}

type Transaction struct {
    ID          int             `json:"id"`
	UserID	    int             `json:"userID"`
    AccountID   int             `json:"accountID"`                // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
    Date        time.Time       `json:"date"`
    Payee       string          `json:"payee"`
    Amount      float64         `json:"amount"`
    Memo        string          `json:"memo"`
    CategoryID  int             `json:"categoryID"`                // Simplified version, consider interface{} or struct{} for complex scenarios
}

type Account struct {
    ID       int                `json:"id"`
	UserID	 int                `json:"userID"`
    Name     string             `json:"name"`
    Type     AccountType        `json:"type"`
    BankName string             `json:"bankName"`
    Balance  float64            `json:"balance"`
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