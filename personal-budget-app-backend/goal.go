package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Goal struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Amount      int         `json:"amount"` // in cents not dollars
	TargetDate  time.Time   `json:"targetDate"`
	CategoryID  int         `json:"categoryID"`
	Periodicity Periodicity `json:"periodicity"`
}

type Periodicity string

const (
	Onetime   Periodicity = "onetime"
	Weekly    Periodicity = "weekly"
	BiWeekly  Periodicity = "biweekly"
	Monthly   Periodicity = "monthly"
	Quarterly Periodicity = "quarterly"
	Yearly    Periodicity = "yearly"
)

func (g *Goal) create() error {
	db := initializeDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO goals (email, name, amount, target_date, category_id, periodicity) VALUES (?, ?, ?, ?, ?, ?)", g.Email, g.Name, g.Amount, g.TargetDate, g.CategoryID, g.Periodicity)
	if err != nil {
		return err
	}
	return nil
}

func getGoalByCategoryID(c *gin.Context) {
	var goal Goal
	// get goal
	categoryID := c.Param("categoryID")
	fmt.Println("categoryID: ", categoryID)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM goals WHERE category_id = ?", categoryID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goal"})
		return
	}
	for rows.Next() {
		var tempDate []uint8
		err := rows.Scan(&goal.ID, &goal.Email, &goal.Name, &goal.Amount, &tempDate, &goal.CategoryID, &goal.Periodicity)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goal"})
			return
		} else {
			goal.TargetDate, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing goal target date"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, goal)
}

func getGoalsByEmail(c *gin.Context) {
	var goal Goal
	// get goals
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := initializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM goals WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
		return
	}
	var goals []Goal
	for rows.Next() {
		var tempDate []uint8
		err := rows.Scan(&goal.ID, &goal.Email, &goal.Name, &goal.Amount, &tempDate, &goal.CategoryID, &goal.Periodicity)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
			return
		} else {
			goal.TargetDate, err = time.Parse("2006-01-02T00:00:00Z04:00", string(tempDate))
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing goal target date"})
				return
			}
			goals = append(goals, goal)
		}
	}
	c.JSON(http.StatusOK, goals)
}

func postGoal(c *gin.Context) {
	var newGoal Goal
	if err := c.BindJSON(&newGoal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newGoal)
	if err := newGoal.create(); err != nil {
		fmt.Println("error in newGoal.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newGoal)
}
