package models

import (
	"personal-budget-app-backend/database"
)

type MonthlyBudget struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	CategoryID int    `json:"categoryID"`
	Assigned   int    `json:"assigned"` // in cents not dollars
	Spent      int    `json:"spent"`    // in cents not dollars
	Balance    int    `json:"balance"`  // in cents not dollars, not in DB. calculated as assigned - spent
}

func (m *MonthlyBudget) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	// check to see if monthly budget already exists
	var count int
	var query string
	err := db.QueryRow("SELECT COUNT(*) FROM monthly_budgets WHERE email = ? AND month = ? AND year = ? AND category_id = ?", m.Email, m.Month, m.Year, m.CategoryID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		query = "UPDATE monthly_budgets SET assigned = ? WHERE email = ? AND month = ? AND year = ? AND category_id = ?"
		_, err := db.Exec(query, m.Assigned, m.Email, m.Month, m.Year, m.CategoryID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err = db.Exec("INSERT INTO monthly_budgets (email, month, year, category_id, assigned, spent) VALUES (?, ?, ?, ?, ?, 0)", m.Email, m.Month, m.Year, m.CategoryID, m.Assigned)
	if err != nil {
		return err
	}
	return nil
}

func GetMonthlyBudgets(email string, month int, year int) ([]MonthlyBudget, error) {
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM monthly_budgets WHERE email = ? AND month = ? AND year = ?", email, month, year)
	if err != nil {
		return nil, err
	}
	var monthlyBudgets []MonthlyBudget
	for rows.Next() {
		var monthlyBudget MonthlyBudget
		err := rows.Scan(&monthlyBudget.ID, &monthlyBudget.Email, &monthlyBudget.Month, &monthlyBudget.Year, &monthlyBudget.CategoryID, &monthlyBudget.Assigned, &monthlyBudget.Spent)
		if err != nil {
			return nil, err
		}
		monthlyBudget.Balance = monthlyBudget.Assigned - monthlyBudget.Spent
		monthlyBudgets = append(monthlyBudgets, monthlyBudget)
	}
	return monthlyBudgets, nil
}
