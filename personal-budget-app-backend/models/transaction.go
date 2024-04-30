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
	// check for existence of monthly_budget for this category. if results are null, insert new row
	row, err := db.Query("SELECT * from monthly_budgets where category_id = ? AND month = ? AND year = ?", t.CategoryID, t.Date.Month(), t.Date.Year())
	if err != nil {
		return err
	}
	if !row.Next() {
		_, err = db.Exec("INSERT INTO monthly_budgets (email, category_id, month, year, assigned, spent) VALUES (?, ?, ?, ?, 0, ?)", t.Email, t.CategoryID, t.Date.Month(), t.Date.Year(), -t.Amount)
		if err != nil {
			return err
		}
	} else {
		if err := t.updateMonthlyBudgetSpent(); err != nil {
			return err
		}
	}
	if err := t.updateAccountBalance(); err != nil {
		return err
	}
	fmt.Println("Transaction created and account balance and monthly budget updated for account: ", t.AccountID)
	return nil
}

func GetTransactionsByEmail(email string) ([]Transaction, error) {
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query(`
	SELECT t.id, t.account_id, DATE_FORMAT(t.date, '%Y-%m-%d %H:%i:%s'), t.payee, t.amount, t.memo, t.category_id, t.email, c.name 
	FROM transactions t
	LEFT JOIN categories c ON c.id = t.category_id 
	WHERE t.email = ?`, email)
	if err != nil {
		return nil, err
	}
	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		var tempDate []uint8
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &tempDate, &transaction.Payee, &transaction.Amount, &transaction.Memo, &transaction.CategoryID, &transaction.Email, &transaction.CategoryName)
		if err != nil {
			return nil, err
		}
		fmt.Println("tempDate", tempDate)
		fmt.Println("tempDate", string(tempDate))
		transaction.Date, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t *Transaction) updateAccountBalance() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_account_balance(?, ?)", t.AccountID, t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) updateMonthlyBudgetSpent() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", t.CategoryID, t.Date.Month(), t.Date.Year(), -t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) reverseUpdateAccountBalance() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_account_balance(?, ?)", t.AccountID, -t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) reverseUpdateMonthlyBudgetSpent() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", t.CategoryID, t.Date.Month(), t.Date.Year(), -t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) reverseUpdateAccountAndMonthlyBudget() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("CALL update_account_balance(?, ?)", t.AccountID, -t.Amount)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL update_monthly_budget_spent(?, ?, ?, ?)", t.CategoryID, t.Date.Month(), t.Date.Year(), t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (ot *Transaction) Update(nt *Transaction) error {
	// method called on old transaction (ot) with param new transaction (nt)
	db := database.InitializeDB()
	defer db.Close()
	// reverse old transaction updates
	err := ot.reverseUpdateAccountBalance()
	if err != nil {
		return err
	}
	err = ot.reverseUpdateMonthlyBudgetSpent()
	if err != nil {
		return err
	}
	// update transaction
	_, err = db.Exec("UPDATE transactions SET account_id = ?, date = ?, payee = ?, amount = ?, memo = ?, category_id = ?, email = ? WHERE id = ?", nt.AccountID, nt.Date, nt.Payee, nt.Amount, nt.Memo, nt.CategoryID, nt.Email, nt.ID)
	if err != nil {
		return err
	}
	err = nt.updateAccountBalance()
	if err != nil {
		return err
	}
	err = nt.updateMonthlyBudgetSpent()
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Delete() error {
	db := database.InitializeDB()
	defer db.Close()
	// get transaction amount
	var amount int
	err := db.QueryRow("SELECT amount FROM transactions WHERE id = ?", t.ID).Scan(&amount)
	if err != nil {
		return err
	}
	// reverse transaction updates
	err = t.reverseUpdateAccountAndMonthlyBudget()
	if err != nil {
		return err
	}
	// delete transaction
	_, err = db.Exec("DELETE FROM transactions WHERE id = ?", t.ID)
	if err != nil {
		return err
	}
	fmt.Println("Transaction deleted and account balance and monthly budget updated for account: ", t.AccountID)
	return nil
}
