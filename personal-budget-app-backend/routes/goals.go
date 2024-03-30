package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterGoalsRoutes(router *gin.Engine) error {
	router.GET("/goals/:email/:categoryID/:month/:year", GetGoals)
	router.GET("/goals/category/:categoryID", GetGoalsByCategoryID)
	router.POST("/goals", PostGoal)
	return nil
}

func GetGoals(c *gin.Context) {
	// get goals
	email := c.Param("email")
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	goals, err := models.GetGoals(email, categoryID, month, year)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, goals)
}

func GetGoalsByCategoryID(c *gin.Context) {
	var goal models.Goal
	// get goal
	categoryID := c.Param("categoryID")
	fmt.Println("categoryID: ", categoryID)
	db := database.InitializeDB()
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

func GetGoalsByEmail(c *gin.Context) {
	var goal models.Goal
	// get goals
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM goals WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
		return
	}
	var goals []models.Goal
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

func PostGoal(c *gin.Context) {
	var newGoal models.Goal
	if err := c.BindJSON(&newGoal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newGoal)
	if err := newGoal.Save(); err != nil {
		fmt.Println("error in newGoal.create(): ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newGoal)
}
