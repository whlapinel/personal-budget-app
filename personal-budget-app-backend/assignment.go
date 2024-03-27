package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Assignment struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	CategoryID int    `json:"categoryID"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Amount     int    `json:"amount"` // in cents not dollars
}

func (a *Assignment) Save() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO assignments (email, category_id, month, year, amount) VALUES (?, ?, ?, ?, ?)", a.Email, a.CategoryID, a.Month, a.Year, a.Amount)
	if err != nil {
		return err
	}
	return nil
}

func getAssignmentsByCategoryID(c *gin.Context) {
	var assignment Assignment
	// get assignments
	categoryID := c.Param("categoryID")
	fmt.Println("categoryID: ", categoryID)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM assignments WHERE category_id = ?", categoryID)
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
	// see if there's already an assignment for this category, month, and year
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM assignments WHERE category_id = ? AND month = ? AND year = ?", newAssignment.CategoryID, newAssignment.Month, newAssignment.Year)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting assignments"})
		return
	}
	for rows.Next() {
		var assignment Assignment
		err := rows.Scan(&assignment.ID, &assignment.Email, &assignment.CategoryID, &assignment.Month, &assignment.Year, &assignment.Amount)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting assignments"})
			return
		} else {
			// overwrite the assignment amount
			assignment.Amount = newAssignment.Amount
			// update the assignment row in DB
			_, err := db.Exec("UPDATE assignments SET amount = ? WHERE id = ?", assignment.Amount, assignment.ID)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating assignment"})
				return
			}
			c.JSON(http.StatusOK, assignment)
			return
		}
	}
	// no assignment found, so create a new one
	fmt.Println(newAssignment)
	if err := newAssignment.Save(); err != nil {
		fmt.Println("error in newAssignment.Save(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAssignment)
}
