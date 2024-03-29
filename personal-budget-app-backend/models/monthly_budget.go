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
	Amount     int    `json:"amount"` // in cents not dollars
}

func (m *MonthlyBudget) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO monthly_budgets (email, month, year, category_id, amount) VALUES (?, ?, ?, ?, ?)", m.Email, m.Month, m.Year, m.CategoryID, m.Amount)
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
		err := rows.Scan(&monthlyBudget.ID, &monthlyBudget.Email, &monthlyBudget.Month, &monthlyBudget.Year, &monthlyBudget.CategoryID, &monthlyBudget.Amount)
		if err != nil {
			return nil, err
		}
		monthlyBudgets = append(monthlyBudgets, monthlyBudget)
	}
	return monthlyBudgets, nil
}
