package viewmodels

import (
	"personal-budget-app-backend/database"
	"time"
)

type ReportEntry struct {
	Email    string `json:"email"`
	Month    int    `json:"month"`
	Year     int    `json:"year"`
	Income   int    `json:"income"`
	Expenses int    `json:"expenses"`
}

// gets current and last 3 months income & expenses for reports page
func GetIncomeAndExpenses(email string) (*[]ReportEntry, error) {
	// get current month
	today := time.Now()
	thisMonth := int(today.Month())
	thisYear := today.Year()
	db := database.InitializeDB()
	defer db.Close()
	var incomeExpenses []ReportEntry
	for i := 2; i >= 0; i-- {
		currMonth := thisMonth - i
		currYear := thisYear
		if currMonth <= 0 {
			currMonth = 12 + currMonth
			currYear = thisYear - 1
		}
		var income int
		err := db.QueryRow("SELECT IFNULL(SUM(amount), 0) FROM transactions WHERE email = ? AND month(date) = ? AND year(date) = ? AND amount > 0", email, currMonth, currYear).Scan(&income)
		if err != nil {
			return nil, err
		}
		// get expenses
		var expenses int
		err = db.QueryRow("SELECT IFNULL(SUM(amount) * -1, 0) FROM transactions WHERE email = ? AND month(date) = ? AND year(date) = ? AND amount < 0", email, currMonth, currYear).Scan(&expenses)
		if err != nil {
			return nil, err
		}
		// return IncomeExpenseReport
		incomeExpenses = append(incomeExpenses, ReportEntry{
			Email:    email,
			Month:    currMonth,
			Year:     currYear,
			Income:   income,
			Expenses: expenses,
		})
	}
	return &incomeExpenses, nil
}
