package models

import (
	"personal-budget-app-backend/database"
)

type Assignment struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	CategoryID int    `json:"categoryID"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Amount     int    `json:"amount"` // in cents not dollars
}



func (a *Assignment) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO assignments (email, category_id, month, year, amount) VALUES (?, ?, ?, ?, ?)", a.Email, a.CategoryID, a.Month, a.Year, a.Amount)
	if err != nil {
		return err
	}
	return nil
}
