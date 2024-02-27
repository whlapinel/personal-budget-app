package main


import (
	"database/sql"
	"fmt"
)


func seedUserTestData(db *sql.DB) (sql.Result, error) {
	var result sql.Result
	for i, v := range users {
		fmt.Println("Inserting user", i, "into database")
		result, err := db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)", v.FirstName, v.LastName, v.Email)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
	return result, nil
}

func seedCategoryTestData(db *sql.DB) (sql.Result, error) {
	var result sql.Result
	for i, v := range budgetCategories {
		fmt.Println("Inserting category", i, "into database")
		result, err := db.Exec("INSERT INTO categories (user_id, name, needed, assigned, spent, available) VALUES (?, ?, ?, ?, ?, ?)", v.UserID, v.Name, v.Needed, v.Assigned, v.Spent, v.Available)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
	return result, nil
}
