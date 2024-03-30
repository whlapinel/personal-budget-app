package models

import (
	"personal-budget-app-backend/database"
)

type ProjIncome struct {
	ID          int         `json:"id"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Amount      int         `json:"amount"` // in cents not dollars
	TargetDate  string      `json:"targetDate"`
	Periodicity Periodicity `json:"periodicity"` // onetime or monthly
}

func (pi *ProjIncome) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO proj_incomes (email, name, amount, target_date, periodicity) VALUES (?, ?, ?, ?, ?)", pi.Email, pi.Name, pi.Amount, pi.TargetDate, pi.Periodicity)
	if err != nil {
		return err
	}
	return nil
}
