package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Assignment struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	CategoryID int    `json:"categoryID"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Amount     int    `json:"amount"` // in cents not dollars
}

func (a *Assignment) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO assignments (email, category_id, month, year, amount) VALUES (?, ?, ?, ?, ?)", a.Email, a.CategoryID, a.Month, a.Year, a.Amount)
	if err != nil {
		return err
	}
	return nil
}

func getAssignmentsByEmail(c *gin.Context) {
	var assignment Assignment
	// get assignments
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM assignments WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting assignments"})
		return
	}
	var assignments []Assignment
	for rows.Next() {
		err := rows.Scan(&assignment.ID, &assignment.Email, &assignment.CategoryID, &assignment.Month, &assignment.Year, &assignment.Amount)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting assignments"})
			return
		} else {
			assignments = append(assignments, assignment)
		}
	}
	c.JSON(http.StatusOK, assignments)
}

func postAssignment(c *gin.Context) {
	var newAssignment Assignment
	if err := c.BindJSON(&newAssignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newAssignment)
	if err := newAssignment.create(); err != nil {
		fmt.Println("error in newAssignment.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAssignment)
}
