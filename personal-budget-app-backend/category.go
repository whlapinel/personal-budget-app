package main

import (
	"fmt"
)

type Category struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

func (bc *Category) create() error {
	fmt.Println("Creating category")
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO categories (email, name) VALUES (?, ?)", bc.Email, bc.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

