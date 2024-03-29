package models

import (
	"fmt"
	"personal-budget-app-backend/database"
	"time"
)

type Transaction struct {
	ID           int       `json:"id"`
	AccountID    int       `json:"accountID"` // Changed from Account['id'] to string to simplify, consider interface{} if needing more complexity
	Date         time.Time `json:"date"`
	Payee        string    `json:"payee"`
	Amount       int       `json:"amount"`       // in cents not dollars
	Memo         *string   `json:"memo"`         // pointer so value can be nil
	CategoryID   *int      `json:"categoryID"`   // pointer so value can be nil
	CategoryName *string   `json:"categoryName"` // stored in DB under categories.name
	Email        string    `json:"email"`
}


func (t *Transaction) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO transactions (account_id, date, payee, amount, memo, category_id, email) VALUES (?, ?, ?, ?, ?, ?, ?)", t.AccountID, t.Date, t.Payee, t.Amount, t.Memo, t.CategoryID, t.Email)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_account_balance(?, ?)", t.AccountID, t.Amount)
	if err != nil {
		return err
	}
	fmt.Println("Transaction created and account balance updated for account: ", t.AccountID)
	return nil
}
