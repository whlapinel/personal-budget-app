package models

import (
	"fmt"
	"personal-budget-app-backend/database"
)

type Category struct {
	ID      int     `json:"id"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Balance int     `json:"balance"`
	Goals   *[]Goal `json:"goals"` // not stored in DB, but should be retrieved along with category
}


func (bc *Category) Save() error {
	fmt.Println("Creating category")
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO categories (email, name) VALUES (?, ?)", bc.Email, bc.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (bc *Category) Get(id int) (*Category, error) {
	fmt.Println("Getting category by ID")
	db := database.InitializeDB()
	defer db.Close()
	err := db.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(&bc.ID, &bc.Email, &bc.Name, &bc.Balance)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bc, nil
}


func GetCategories(email string) ([]Category, error) {
	fmt.Println("running getCategories()")
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT id, email, name, IFNULL(balance, 0) FROM categories WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Email, &category.Name, &category.Balance)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
