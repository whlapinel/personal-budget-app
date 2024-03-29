package models

import (
	"personal-budget-app-backend/database"
	"time"
)

type Goal struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Amount      int         `json:"amount"` // in cents not dollars
	TargetDate  time.Time   `json:"targetDate"`
	CategoryID  int         `json:"categoryID"`
	Periodicity Periodicity `json:"periodicity"`
}

type Periodicity string

const (
	Onetime   Periodicity = "onetime"
	Monthly   Periodicity = "monthly"
)



func (g *Goal) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO goals (email, name, amount, target_date, category_id, periodicity) VALUES (?, ?, ?, ?, ?, ?)", g.Email, g.Name, g.Amount, g.TargetDate, g.CategoryID, g.Periodicity)
	if err != nil {
		return err
	}
	return nil
}


func GetGoals(categoryID string) ([]Goal, error) {
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM goals WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	var goals []Goal
	for rows.Next() {
		var goal Goal
		var tempDate []uint8
		err := rows.Scan(&goal.ID, &goal.Email, &goal.Name, &goal.Amount, &tempDate, &goal.CategoryID, &goal.Periodicity)
		if err != nil {
			return nil, err
		}
		goal.TargetDate, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}
	return goals, nil
}
