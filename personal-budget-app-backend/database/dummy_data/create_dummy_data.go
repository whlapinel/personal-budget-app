package dummy_data

import (
	"fmt"
	"os"
	"personal-budget-app-backend/models"
	"time"
)

func CreateDummyData() error {
	// _, err := ExecuteScript("add_dummy_data.sql", ";\n", "exec")
	// return err
	test_user := models.User{
		Email:     "test@test.com",
		Password:  os.Getenv("TEST_USER_PASSWORD"),
		FirstName: "Test",
		LastName:  "User",
	}
	if err := test_user.Save(); err != nil {
		return err
	}
	test_account := models.Account{
		Email:           "test@test.com",
		Name:            "Test Account",
		Type:            "Checking",
		BankName:        "Test Bank",
		StartingBalance: 1000000,
		Balance:         1000000,
	}
	if err := test_account.Save(); err != nil {
		return err
	}
	createCategory := func(name string) error {
		category := models.Category{
			Email: "test@test.com",
			Name:  name,
		}
		if err := category.Save(); err != nil {
			return err
		}
		return nil
	}
	cat_names := []string{"Utilities", "Groceries", "Entertainment", "Auto", "Vacation"}
	for _, name := range cat_names {
		if err := createCategory(name); err != nil {
			return err
		}
	}
	getTime := func(date string) time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05Z", date+"T00:00:00Z")
		return t
	}
	testDate := getTime("2024-05-01")
	fmt.Println("CreateDummyData testDate: ", testDate)
	goals := []models.Goal{
		{
			Email:       "test@test.com",
			Name:        "Electric Bill",
			Amount:      30000,
			TargetDate:  getTime("2024-5-25"),
			CategoryID:  1,
			Periodicity: "monthly",
		},
		{
			Email:       "test@test.com",
			Name:        "Internet",
			Amount:      10000,
			TargetDate:  getTime("2025-12-25"),
			CategoryID:  1,
			Periodicity: "monthly",
		},
	}
	for _, goal := range goals {
		if err := goal.Save(); err != nil {
			return err
		}
	}
	monthly_budgets := []models.MonthlyBudget{
		{
			Email:      "test@test.com",
			Month:      5,
			Year:       2024,
			CategoryID: 1,
			Assigned:   30000,
		},
	}
	for _, monthly_budget := range monthly_budgets {
		if err := monthly_budget.Save(); err != nil {
			return err
		}
	}
	fmt.Println("getTime() test: ", getTime("2024-5-1"))

	transactions := []models.Transaction{
		{
			Email:      "test@test.com",
			Date:       getTime("2024-05-01"),
			Amount:     -30000,
			AccountID:  1,
			CategoryID: &[]int{1}[0],
		},
	}
	fmt.Println("transaction date: ", transactions[0].Date)
	for _, transaction := range transactions {
		if err := transaction.Save(); err != nil {
			return err
		}
	}
	return nil
}
