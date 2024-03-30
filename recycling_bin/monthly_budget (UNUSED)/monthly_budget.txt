package monthly_budget

import (
	"fmt"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
)

type MonthlyBudget struct {
	Email                     string        `json:"email"`
	Month                     int           `json:"month"`
	Year                      int           `json:"year"`
	BudgetFunds               int           `json:"budgetFunds"`
	Unassigned                int           `json:"unassigned"`
	TotalAssignedCurrentMonth int           `json:"totalAssignedCurrentMonth"`
	AssignedInFuture          int           `json:"assignedInFuture"`
	CategoryArray             CategoryArray `json:"categoryArray"`
}

type CategoryArray []CategoryData

type CategoryData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Needed    int    `json:"needed"`
	Assigned  int    `json:"assigned"`
	Available int    `json:"available"`
	Spent     int    `json:"spent"`
}

func GetMonthlyBudget(email string, month int, year int) (*MonthlyBudget, error) {
	fmt.Println("running getMonthlyBudget()")
	var monthlyBudget MonthlyBudget
	monthlyBudget.Email = email
	monthlyBudget.Month = month
	monthlyBudget.Year = year
	categoryMap, err := getCategoryMap(email, month, year)
	if err != nil {
		return nil, err
	}
	monthlyBudget.CategoryArray = *categoryMap
	budgetFunds, err := getBudgetFunds(email)
	if err != nil {
		return nil, err
	}
	monthlyBudget.BudgetFunds = budgetFunds

	assignedThisMonth, err := getAssignedCurrentMonth(email, month, year)
	if err != nil {
		return nil, err
	}
	monthlyBudget.TotalAssignedCurrentMonth = assignedThisMonth
	assignedInFuture, err := getAssignedInFuture(email, month, year)
	if err != nil {
		return nil, err
	}
	monthlyBudget.AssignedInFuture = assignedInFuture


	availableLastMonth := getAvailableLastMonth(email, month, year)


	monthlyBudget.Unassigned = budgetFunds - assignedThisMonth - assignedInFuture - availableLastMonth

	return &monthlyBudget, nil
}


func getAssignedTotal(email string, month int, year int) (int, error) {
	fmt.Println("running getAssignedTotal()")
	db := database.InitializeDB()
	defer db.Close()
	assignedTotalQuery := `
	SELECT IFNULL(SUM(amount),0)
	FROM assignments
	WHERE email = ? AND month = ? AND year = ?
	`
	rows, err := db.Query(assignedTotalQuery, email, month, year)
	if err != nil {
		return 0, err
	}
	var assignedTotal int
	for rows.Next() {
		err := rows.Scan(&assignedTotal)
		if err != nil {
			return 0, nil
		}
	}
	return assignedTotal, nil
}

func getSpentTotal(email string, month int, year int) (int, error) {
	fmt.Println("running getSpentTotal()")
	db := database.InitializeDB()
	defer db.Close()
	spentTotalQuery := `
	SELECT IFNULL(SUM(amount) * -1,0)
	FROM transactions
	WHERE email = ? AND month(date) = ? AND year(date) = ?
	`
	rows, err := db.Query(spentTotalQuery, email, month, year)
	if err != nil {
		return 0, err
	}
	var spentTotal int
	for rows.Next() {
		err := rows.Scan(&spentTotal)
		if err != nil {
			return 0, nil
		}
	}
	return spentTotal, nil
}

func getAvailableLastMonth(email string, month int, year int) int {
	fmt.Println("running getAvailableLastMonth()")
	var lastMonth int
	var yearOfPrevMonth int
	if month == 0 {
		lastMonth = 11
		yearOfPrevMonth = year - 1
	} else {
		lastMonth = month - 1
		yearOfPrevMonth = year
	}
	assignedLastMonth, err := getAssignedTotal(email, lastMonth, yearOfPrevMonth)
	if err != nil {
		return 0
	}
	spentLastMonth, err := getSpentTotal(email, lastMonth, yearOfPrevMonth)
	if err != nil {
		return 0
	}
	availableLastMonth := assignedLastMonth - spentLastMonth
	return availableLastMonth
}

func getAssignedInFuture(email string, month int, year int) (int, error) {
	// get all assignments for all future months
	fmt.Println("running getAssignedInFuture()")
	db := database.InitializeDB()
	defer db.Close()
	assignedInFutureQuery := `
	SELECT IFNULL(SUM(amount),0)
	FROM assignments
	WHERE email = ? AND month > ? AND year >= ?
	`
	rows, err := db.Query(assignedInFutureQuery, email, month, year)
	if err != nil {
		return 0, err
	}
	var assignedInFuture int
	for rows.Next() {
		err := rows.Scan(&assignedInFuture)
		if err != nil {
			return 0, nil
		}
	}
	return assignedInFuture, nil
}

func getAssignedCurrentMonth(email string, month int, year int) (int, error) {
	fmt.Println("running getAssignedThisMonth()")
	var assignedThisMonth int
	db := database.InitializeDB()
	defer db.Close()
	assignedThisMonthQuery := `
	SELECT IFNULL(SUM(amount),0)
	FROM assignments
	WHERE email = ? AND month = ? AND year = ?
	`
	rows, err := db.Query(assignedThisMonthQuery, email, month, year)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err := rows.Scan(&assignedThisMonth)
		if err != nil {
			return 0, nil
		}
	}
	return assignedThisMonth, nil
}

func getBudgetFunds(email string) (int, error) {
	fmt.Println("running getBudgetFunds()")
	var budgetFunds int
	db := database.InitializeDB()
	defer db.Close()
	budgetFundsQuery :=
		`
	SELECT SUM(balance)
	FROM accounts
	WHERE email = ?	
	`
	rows, err := db.Query(budgetFundsQuery, email)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err := rows.Scan(&budgetFunds)
		if err != nil {
			return 0, nil
		}
	}
	return budgetFunds, nil
}

func getCategoryData(categoryID int, month int, year int) (*CategoryData, error) {
	fmt.Println("running getCategoryData")
	var categoryData CategoryData
	// get needed = total goals for category and month
	needed, err := getNeeded(categoryID, month, year)
	if err != nil {
		return nil, err
	}
	categoryData.Needed = needed
	// get assigned = total assignments for category and month
	assigned, err := getAssigned(categoryID, month, year)
	if err != nil {
		return nil, err
	}
	categoryData.Assigned = assigned
	// get spent = total transactions for category and month
	spent, err := getSpent(categoryID, month, year)
	if err != nil {
		return nil, err
	}
	categoryData.Spent = spent
	// get available = assigned for category and month, minus spent, plus any available from previous month
	var lastMonth int
	var yearOfPrevMonth int
	if month == 0 {
		lastMonth = 11
		yearOfPrevMonth = year - 1
	} else {
		lastMonth = month - 1
		yearOfPrevMonth = year
	}
	assignedLastMonth, err := getAssigned(categoryID, lastMonth, yearOfPrevMonth)
	if err != nil {
		return nil, err
	}
	spentLastMonth, err := getSpent(categoryID, lastMonth, yearOfPrevMonth)
	if err != nil {
		return nil, err
	}
	availableLastMonth := assignedLastMonth - spentLastMonth
	categoryData.Available = assigned - spent + availableLastMonth

	return &categoryData, nil
}

func getCategoryMap(email string, month int, year int) (*CategoryArray, error) {
	fmt.Println("running getCategoryMap()")
	var categoryArray CategoryArray
	categories, err := models.GetCategories(email)
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		categoryData, err := getCategoryData(category.ID, month, year)
		if err != nil {
			return nil, err
		}
		categoryData.ID = category.ID
		categoryData.Name = category.Name
		// append categoryData to categoryArray
		categoryArray = append(categoryArray, *categoryData)
	}
	return &categoryArray, nil
}


func getSpent(categoryID int, month int, year int) (int, error) {
	fmt.Println("running getSpent()")
	db := database.InitializeDB()
	defer db.Close()
	const spentQuery = `SELECT IFNULL(SUM(amount) * -1,0) FROM transactions WHERE category_id = ? AND month(date) = (? + 1) AND year(date) = ?`
	rows, err := db.Query(spentQuery, categoryID, month, year)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var spent int
	for rows.Next() {
		err := rows.Scan(&spent)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}
	return spent, nil
}

func getAssigned(categoryID int, month int, year int) (int, error) {
	fmt.Println("running getAssigned()")
	db := database.InitializeDB()
	defer db.Close()
	const assignedQuery = `SELECT IFNULL(SUM(amount),0) FROM monthly_budget WHERE category_id = ? AND month = ? and year = ?`
	rows, err := db.Query(assignedQuery, categoryID, month, year)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var assigned int
	for rows.Next() {
		err := rows.Scan(&assigned)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}
	return assigned, nil
}

func getNeeded(categoryID int, month int, year int) (int, error) {
	fmt.Println("running getNeeded()")
	db := database.InitializeDB()
	defer db.Close()
	const neededMonthlyQuery = `SELECT IFNULL(SUM(amount),0) FROM goals WHERE category_id = ? AND periodicity = 'monthly'`
	rows, err := db.Query(neededMonthlyQuery, categoryID)
	if err != nil {
		return 0, err
	}
	var neededMonthly int
	for rows.Next() {
		err := rows.Scan(&neededMonthly)
		if err != nil {
			return 0, err
		}
	}
	const onetimeNeededQuery = `SELECT IFNULL(SUM(amount),0) AS one_time_needed FROM goals WHERE periodicity = 'onetime' AND month(target_date) = ? AND year(target_date) = ?`
	rows, err = db.Query(onetimeNeededQuery, month, year)
	if err != nil {
		return 0, err
	}
	var oneTimeNeeded int
	for rows.Next() {
		err := rows.Scan(&oneTimeNeeded)
		if err != nil {
			return 0, err
		}
	}
	needed := neededMonthly + oneTimeNeeded
	return needed, nil
}

