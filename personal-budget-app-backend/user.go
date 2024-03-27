package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

type User struct {
    Email        string     `json:"email"`
    Password     string     `json:"password"`
    FirstName    string     `json:"firstName"`
    LastName     string     `json:"lastName"`
}

func (u *User) Save() error {
	db := initializeDB()
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

func getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	fmt.Println(email)
	db := initializeDB()
	defer db.Close()
	fmt.Println("db initialized")
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}


func postUser(c *gin.Context) {
	var user *User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
