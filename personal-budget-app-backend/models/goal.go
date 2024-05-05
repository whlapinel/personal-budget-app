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
	Onetime Periodicity = "onetime"
	Monthly Periodicity = "monthly"
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

func GetGoals(email string, categoryID, month, year int) (*[]Goal, error) {
	var goal Goal
	db := database.InitializeDB()
	defer db.Close()
	// this should cover both:
	// - monthly goals with dates whose month falls
	// within current or past months,
	// and:
	// - onetime goals with a target_date that falls
	// in the future, divided by the number of months until then
	rows, err := db.Query(`
	SELECT *
	FROM goals
	WHERE email = ?
  		AND category_id = ?
  		AND (
    	  (periodicity = 'monthly' AND DATE_FORMAT(target_date, '%Y-%m') <= DATE_FORMAT(STR_TO_DATE(CONCAT(?,'-',?,'-01'), '%Y-%m-%d'), '%Y-%m'))
    	  OR
    	  (periodicity = 'onetime' AND DATE_FORMAT(target_date, '%Y-%m') > DATE_FORMAT(STR_TO_DATE(CONCAT(?,'-',?,'-01'), '%Y-%m-%d'), '%Y-%m'))
  		);
`,
		email, categoryID, month, year)
	if err != nil {
		return nil, err
	}
	var goals []Goal
	for rows.Next() {
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
	return &goals, nil
}
