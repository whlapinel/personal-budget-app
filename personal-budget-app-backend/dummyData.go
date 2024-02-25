package main

var users = []User {
	{
		ID:        "user1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	},
}

var budgetCategories = []BudgetCategory {
	{
		ID:        "category1",
		UserID:    "user1",
		Name:      "Groceries",
		Needed:    300.00,
		Assigned:  300.00,
		Spent:     0.00,
		Available: 300.00,	
	},
}