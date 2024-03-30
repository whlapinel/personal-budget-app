package viewmodels

import (
	"fmt"
	"personal-budget-app-backend/database"
)

type BudgetPageData struct {
	CurrentFunds    int            `json:"currentFunds"`    // sum of all account balances
	ProjIncome      int            `json:"projIncome"`      // sum of all projected incomes for month
	TotalAvailable  int            `json:"totalAvailable"`  // sum of balances for all monthly budgets
	TotalUnassigned int            `json:"totalUnassigned"` // current funds minus sum of all balances for all monthly monthly budgets
	CategoryRows    []CategoryData `json:"categoryRows"`
}

type CategoryData struct {
	CategoryID   int    `json:"categoryID"`   // from category
	CategoryName string `json:"categoryName"` // from category
	Assigned     int    `json:"assigned"`     // from monthly budget
	Spent        int    `json:"spent"`        // from monthly budget
	Available    int    `json:"available"`    // from monthly budget, includes balances of past months but not future months
	GoalsSum     int    `json:"goalsSum"`     // goals for category
}

func GetBudgetPageData(email string, month, year int) (*BudgetPageData, error) {
	var budgetPageData BudgetPageData
	db := database.InitializeDB()
	defer db.Close()
	// get CategoryID, CategoryName, Assigned, Spent, Available, GoalsSum
	query := `
	SELECT c.id, c.name, IFNULL(m.assigned, 0), IFNULL(m.spent, 0) 
	FROM categories c
	LEFT JOIN monthly_budgets m ON c.id = m.category_id AND m.month = ? AND m.year = ?
	WHERE c.email = ?`
	rows, err := db.Query(query, month, year, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var categoryData CategoryData
		err := rows.Scan(&categoryData.CategoryID, &categoryData.CategoryName, &categoryData.Assigned, &categoryData.Spent)
		if err != nil {
			return nil, err
		}
		budgetPageData.CategoryRows = append(budgetPageData.CategoryRows, categoryData)
	}
	fmt.Println("GetBudgetPageData(): budgetPageData.CategoryRows: ", budgetPageData.CategoryRows)
	// get GoalsSum
	query = `
	SELECT IFNULL(SUM(amount),0) 
	FROM goals 
	WHERE email = ? 
	AND category_id = ? 
	AND (periodicity = 'monthly' OR (MONTH(target_date) = (?+1) AND YEAR(target_date) = ?))`
	for i := range budgetPageData.CategoryRows {
		err := db.QueryRow(query, email, budgetPageData.CategoryRows[i].CategoryID, month, year).Scan(&budgetPageData.CategoryRows[i].GoalsSum)
		fmt.Println("GetBudgetPageData(): current c_id and goalsSum: ", budgetPageData.CategoryRows[i].CategoryID, budgetPageData.CategoryRows[i].GoalsSum)
		if err != nil {
			return nil, err
		}
	}
	// get Available = SUM(assigned - spent) for current and previous months for each category
	query = `
	SELECT 
	CASE WHEN (assigned - spent) < 0 
	THEN 0
	ELSE IFNULL(SUM(assigned - spent), 0)
	END AS available
	FROM monthly_budgets
	WHERE email = ? AND category_id = ? AND ((year = ? AND month <= ?) OR (year = ? AND month < ?))`
	for i := range budgetPageData.CategoryRows {
		err := db.QueryRow(query, email, budgetPageData.CategoryRows[i].CategoryID, year, month, year, month).Scan(&budgetPageData.CategoryRows[i].Available)
		if err != nil {
			return nil, err
		}
	}
	// get TotalAssigned = SUM(assigned - spent) for all months
	query = `SELECT 
	CASE WHEN SUM(assigned - spent) < 0
	THEN 0
	ELSE IFNULL(SUM(assigned - spent), 0) 
	END AS total_assigned
	FROM monthly_budgets WHERE email = ?`
	err = db.QueryRow(query, email).Scan(&budgetPageData.TotalAvailable)
	if err != nil {
		return nil, err
	}
	// current funds
	query = `SELECT IFNULL(SUM(balance), 0) FROM accounts WHERE email = ?`
	err = db.QueryRow(query, email).Scan(&budgetPageData.CurrentFunds)
	if err != nil {
		return nil, err
	}
	// get TotalUnassigned = current funds - TotalAvailable
	budgetPageData.TotalUnassigned = budgetPageData.CurrentFunds - budgetPageData.TotalAvailable

	return &budgetPageData, nil
}
