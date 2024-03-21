package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
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
	id := u.Email
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

func postUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newUser)
	if err := newUser.create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
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

