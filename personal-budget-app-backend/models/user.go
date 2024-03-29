package models

import (
	"database/sql"
	"fmt"
	"personal-budget-app-backend/database"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}


func (u *User) Save() error {
	db := database.InitializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO users (password, first_name, last_name, email) VALUES (?, ?, ?, ?)", u.Password, u.FirstName, u.LastName, u.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE email = ?", u.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) Get(db *sql.DB) (*User, error) {
	fmt.Println("Getting user by ID")
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", u.Email).Scan(&u.Email, &u.FirstName, &u.LastName, &u.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return u, nil
}

func (u *User) Update() error {
	fmt.Println("Updating user")
	return nil
}


