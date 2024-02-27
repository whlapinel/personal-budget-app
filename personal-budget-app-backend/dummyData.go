package main

var users = []User{
	{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	},
	{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "janedoe@example.com",
	},
	{
		FirstName: "Jim",
		LastName:  "Doe",
		Email:     "jimdoe@example.com",
	},
}

var budgetCategories = []BudgetCategory{
	{
		UserID:    1,
		Name:      "Groceries",
		Needed:    300.00,
		Assigned:  300.00,
		Spent:     0.00,
		Available: 300.00,
	},
	{
		UserID:    2,
		Name:      "Gas",
		Needed:    100.00,
		Assigned:  100.00,
		Spent:     0.00,
		Available: 100.00,
	},
	{
		UserID:    3,
		Name:      "Groceries",
		Needed:    300.00,
		Assigned:  300.00,
		Spent:     0.00,
		Available: 300.00,
	},
	{
		UserID:    1,
		Name:      "Gas",
		Needed:    100.00,
		Assigned:  100.00,
		Spent:     0.00,
		Available: 100.00,
	},
}
