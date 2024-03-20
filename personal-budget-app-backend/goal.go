package main

import (
	"time"
)

type Goal struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	TargetDate time.Time `json:"targetDate"`
	Category   Category  `json:"category"`
}

func (g *Goal) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO goals (email, name, amount, target_date, category_id) VALUES (?, ?, ?, ?, ?)", g.Email, g.Name, g.Amount, g.TargetDate, g.Category.ID)
	if err != nil {
		return err
	}
	return nil
}
