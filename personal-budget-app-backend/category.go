package main

import (
	"fmt"
)

type Category struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	Name   string `json:"name"`
}

func (bc *Category) create() error {
	fmt.Println("Creating category")
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO categories (user_id, name) VALUES (?, ?)", bc.UserID, bc.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (bc *Category) getByID() *Category {
	fmt.Println("Getting category by ID")
	return bc
}

func (bc *Category) update() error {
	fmt.Println("Updating category")
	return nil
}

func (bc *Category) delete() error {
	id := bc.ID
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Deleted category", id)
	return nil
}

func (bc *Category) getAll() ([]Category, error) {
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories WHERE user_id = ?", bc.UserID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var categories []Category
	var category Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.UserID, &category.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			categories = append(categories, category)
		}
	}
	return categories, nil
}


