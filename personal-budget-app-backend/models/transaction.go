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
	// get month of transaction date
	month := t.Date.Month()
	year := t.Date.Year()

	_, err = db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", t.CategoryID, month, year, t.Amount)
	if err != nil {
		return err
	}
	fmt.Println("Transaction created and account balance and monthly budget updated for account: ", t.AccountID)
	return nil
}

func (t *Transaction) reverseAccountUpdate() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_account_balance(?, ?)", t.AccountID, -t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) reverseMonthlyBudgetUpdate() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", t.CategoryID, t.Date.Month(), t.Date.Year(), -t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (ot *Transaction) Update(nt *Transaction) error {
	// method called on old transaction (ot) with param new transaction (nt)
	db := database.InitializeDB()
	defer db.Close()
	// get previous transaction amount
	err := db.QueryRow("SELECT amount FROM transactions WHERE id = ?", ot.ID).Scan(&ot.Amount)
	if ot.Amount != nt.Amount {
		ot.reverseAccountUpdate()
		ot.reverseMonthlyBudgetUpdate()
	}
	if err != nil {
		return err
	}
	// update transaction
	_, err = db.Exec("UPDATE transactions SET account_id = ?, date = ?, payee = ?, amount = ?, memo = ?, category_id = ?, email = ? WHERE id = ?", nt.AccountID, nt.Date, nt.Payee, nt.Amount, nt.Memo, nt.CategoryID, nt.Email, nt.ID)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_account_balance(?, ?)", nt.AccountID, nt.Amount)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", nt.CategoryID, nt.Date.Month(), nt.Date.Year(), nt.Amount)
	if err != nil {
		return err
	}
	return nil
}
