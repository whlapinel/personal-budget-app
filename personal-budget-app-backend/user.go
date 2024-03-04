package main

import (
	"fmt"
)

type User struct {
    ID           int        `json:"id"`
    Email        string     `json:"email"`
    Password     string     `json:"password"`
    FirstName    string     `json:"firstName"`
    LastName     string     `json:"lastName"`
}

func (u *User) create() error {
    fmt.Println("Creating user")
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO users (password, first_name, last_name, email) VALUES (?, ?, ?, ?)", u.Password, u.FirstName, u.LastName, u.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
    return nil
}

func (u *User) getByID() (*User) {
    fmt.Println("Getting user by ID")
    return u
}

func (u *User) update() error {
    fmt.Println("Updating user")
    return nil
}

func (u *User) delete() error {
	id := u.ID
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
    fmt.Println("Deleted user", id)
    return nil
}
