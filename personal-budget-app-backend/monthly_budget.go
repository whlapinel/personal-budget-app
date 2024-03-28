package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MonthlyBudget struct {
	Email         string        `json:"email"`
	Month         int           `json:"month"`
	Year          int           `json:"year"`
	AssignmentMap AssignmentMap `json:"assignmentMap"`
}

type AssignmentMap map[string]int

func getMonthlyBudgetByEmail(c *gin.Context) {
	var monthlyBudget MonthlyBudget
	// get monthly budget
	email := c.Param("email")
	// coerce month param string into int

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting monthly budget"})
		return
	}		
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting monthly budget"})
		return
	}
	fmt.Println("email: ", email)
	fmt.Println("month: ", month)
	fmt.Println("year: ", year)
	db := initializeDB()
	defer db.Close()
	query := `
	SELECT a.*, c.name 
	FROM assignments a 
	JOIN categories c ON a.category_id = c.id
	WHERE a.email = ? AND month = ? AND year = ?`
	rows, err := db.Query(query, email, month, year)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting monthly budget"})
		return
	}
	assignmentMap := make(AssignmentMap)
	for rows.Next() {
		var assignment Assignment
		var category Category
		err := rows.Scan(&assignment.ID, &assignment.Email, &assignment.CategoryID, &assignment.Month, &assignment.Year, &assignment.Amount, &category.Name)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting monthly budget"})
			return
		} else {
			assignmentMap[category.Name] = assignment.Amount
		}
	}
	monthlyBudget.Email = email
	monthlyBudget.Month = month
	monthlyBudget.Year = year
	monthlyBudget.AssignmentMap = assignmentMap
	c.JSON(http.StatusOK, monthlyBudget)
}