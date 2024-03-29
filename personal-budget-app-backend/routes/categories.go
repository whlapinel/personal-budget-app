package routes

import (
	"fmt"
	"net/http"
	"personal-budget-app-backend/database"
	"personal-budget-app-backend/models"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterCategoriesRoutes(router *gin.Engine) error {
	router.GET("/categories/:email", GetCategoriesByEmailHandler)
	router.POST("/categories", PostCategory)
	return nil
}

func GetCategoriesByEmailHandler(c *gin.Context) {
	var category models.Category
	fmt.Println("running getCategoriesByEmail")
	// get categories
	email := c.Param("email")
	fmt.Println("email: ", email)
	db := database.InitializeDB()
	defer db.Close()
	rows, err := db.Query("SELECT id, email, name, IFNULL(balance, 0) FROM categories WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
		return
	}
	var categories []models.Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Email, &category.Name, &category.Balance)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting categories"})
			return
		} else {
			categories = append(categories, category)
		}
	}
	// get goals for each category
	for i, category := range categories {
		rows, err := db.Query("SELECT * FROM goals WHERE category_id = ?", category.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
			return
		}
		var goals []models.Goal
		for rows.Next() {
			var goal models.Goal
			var tempDate []uint8
			err := rows.Scan(&goal.ID, &goal.Email, &goal.Name, &goal.Amount, &tempDate, &goal.CategoryID, &goal.Periodicity)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting goals"})
				return
			} else {
				goal.TargetDate, err = time.Parse("2006-01-02 00:00:00", string(tempDate))
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "error parsing goal target date"})
					return
				}
				goals = append(goals, goal)
			}
		}
		categories[i].Goals = &goals
	}
	c.JSON(http.StatusOK, categories)
}

func PostCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.BindJSON(&newCategory); err != nil {
		fmt.Println("error in c.BindJSON(&newCategory): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newCategory: ", newCategory)
	if err := newCategory.Save(); err != nil {
		fmt.Println("error in newCategory.Save(): ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCategory)
}
